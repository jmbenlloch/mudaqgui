package main

import (
	"context"
	"fmt"
	"log"
	"math"
	"net"

	"encoding/binary"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type DaqData struct {
	devices                  map[byte]*net.HardwareAddr
	slowControlConfiguration map[byte]map[string]any
	probeConfiguration       map[byte]map[string]any
	events                   map[byte][]EventData
	charges                  map[byte]ChargeHistogram
	chargesRebinned          map[byte]ChargeHistogram
}

type EventData struct {
	//	eventT0    bool
	//	eventT1    bool
	T0         uint32     `json="t0"`
	T1         uint32     `json="t1"`
	LostBuffer uint16     `json="lostBuffer"`
	LostFPGA   uint16     `json="lostFGPA"`
	Charges    [32]uint16 `json="charges"`
}

type ChargeHistogram struct {
	Charges [32][]int32 `json="charges"`
}

func decodeFrame(recvChannel chan Frame, data *DaqData, ctx context.Context) {
	for {
		frame := <-recvChannel
		//log.Printf("length %d, %s", len(frame.Payload), frame.Command)
		switch frame.Command {
		case FEB_OK:
			storeDeviceMac(frame, data, ctx)
			if string(frame.Payload[2:9]) != "FEB_rev" {
				decodeRate(frame, data, ctx)
			}
		case FEB_DATA_CDR:
			log.Println("data cdr")
			decodeData(frame, data, ctx)
		case FEB_EOF_CDR:
			//log.Println("End of data")
		case FEB_OK_SCR:
			log.Println("CITIROC slow control OK")
		case FEB_OK_PMR:
			log.Println("CITIROC probe OK")
		case FEB_OK_FIL:
			log.Println("FPGA input logic OK")
		default:
			log.Fatalf("Unkown response command %s", frame.Command)
		}
	}
}

func storeDeviceMac(frame Frame, data *DaqData, ctx context.Context) {
	//	log.Printf("[%s] %s", frame.Source.String(), string(frame.Payload))
	//	log.Printf("[%s] %x", frame.Source.String(), frame.Payload)
	//	log.Printf("[%s] %x", frame.Source.String(), string(frame.EtherType))
	//	log.Printf("[%s] %x", frame.Source.String(), string(frame.Command))

	data.devices[frame.Source[5]] = &frame.Source
	//fmt.Println(frame.Source[5])
	//fmt.Println(data.devices)
	data.slowControlConfiguration[frame.Source[5]] = createDefaultSlowControlConfiguration()
	data.probeConfiguration[frame.Source[5]] = createDefaultProbeRegisterConfiguration()

	// Add new list to events if this card did not exist
	if _, exists := data.events[frame.Source[5]]; !exists {
		data.events[frame.Source[5]] = make([]EventData, 0, 10000)
	}

	// Add new charge histogram if this card did not exist
	if _, exists := data.charges[frame.Source[5]]; !exists {
		chargeHistograms := ChargeHistogram{}
		for i := 0; i < 32; i++ {
			chargeHistograms.Charges[i] = make([]int32, 1024)
		}
		data.charges[frame.Source[5]] = chargeHistograms

		// Rebinned
		chargeHistogramsRebin := ChargeHistogram{}
		for i := 0; i < 32; i++ {
			chargeHistogramsRebin.Charges[i] = make([]int32, 128)
		}
		data.chargesRebinned[frame.Source[5]] = chargeHistogramsRebin
	}

	cards := make([]int, 0, len(data.devices))
	for key := range data.devices {
		cards = append(cards, int(key))
	}
	// For some reason these values are not properly passed to JS
	//runtime.EventsEmit(ctx, "cards", maps.Keys(data.devices))
	runtime.EventsEmit(ctx, "cards", cards)

	sendConfigToUI(data, ctx)
}

func decodeData(frame Frame, data *DaqData, ctx context.Context) {
	//log.Printf("[%s] %s", frame.Source.String(), string(frame.Payload))
	log.Printf("Payload length [%s] %d", frame.Source.String(), len(frame.Payload))

	// Position 0:2 00 00, unused register value

	data_start := 2
	packet_size := 76

	for data_start < len(frame.Payload)-2 {
		log.Printf("reading: %d - %d Len: %d", data_start, data_start+packet_size, len(frame.Payload))
		evt := decodeEvent(frame.Payload[data_start : data_start+packet_size])
		data_start += packet_size

		data.events[frame.Source[5]] = append(data.events[frame.Source[5]], *evt)

		//log.Printf("[Event lost buffer] %d", evt.LostBuffer)
		//log.Printf("[Event lost fgpa] %d", evt.LostFPGA)
		//log.Printf("[t0] %d", evt.T0)
		//log.Printf("[t1] %d", evt.T1)

		for i := 0; i < 32; i++ {
			log.Printf("charge[%d]: %d", i, evt.Charges[i])
			chargesHistograms := data.charges[frame.Source[5]]
			count := chargesHistograms.Charges[i][evt.Charges[i]]
			chargesHistograms.Charges[i][evt.Charges[i]] = count + 1

			// Rebin
			chargesHistograms = data.chargesRebinned[frame.Source[5]]
			index := evt.Charges[i] / 4
			fmt.Println(index, evt.Charges[i])
			count = chargesHistograms.Charges[i][index]
			chargesHistograms.Charges[i][index] = count + 1
		}
	}

	runtime.EventsEmit(ctx, "events", data.events)
	runtime.EventsEmit(ctx, "charges", data.charges)
	runtime.EventsEmit(ctx, "chargesRebin", data.chargesRebinned)

	//s := hex.EncodeToString(frame.Payload)
	//fmt.Println(s)
	//	log.Printf("[%s] %x", frame.Source.String(), frame.Payload)
	//	log.Printf("[%s] %x", frame.Source.String(), string(frame.EtherType))
	// log.Printf("[%s] %x", frame.Source.String(), string(frame.Command))
}

type CardRate struct {
	Card byte    `json:"card"`
	Rate float32 `json:"rate"`
}

func decodeRate(frame Frame, data *DaqData, ctx context.Context) {
	bits := binary.LittleEndian.Uint32(frame.Payload[2:6])
	rate := math.Float32frombits(bits) // in Hz
	rateEvent := CardRate{Card: frame.Source[5], Rate: rate}
	runtime.EventsEmit(ctx, "rate", rateEvent)
}

func decodeEvent(data []byte) *EventData {
	var eventLostFPGA uint16 = binary.LittleEndian.Uint16(data[0:2])
	var eventLostBuffer uint16 = binary.LittleEndian.Uint16(data[2:4])

	var t0 uint32 = binary.LittleEndian.Uint32(data[4:8])
	var t1 uint32 = binary.LittleEndian.Uint32(data[8:12])

	var t0LSB uint32 = (t0 & 0x00000003)
	var t1LSB uint32 = (t1 & 0x00000003)

	t0 = (t0 & 0x3FFFFFFF) >> 2
	t1 = (t1 & 0x3FFFFFFF) >> 2

	t0 = grayToBin(t0)
	t1 = grayToBin(t1)

	t0 = (t0 << 2) | t0LSB
	t1 = (t1 << 2) | t1LSB

	var charges [32]uint16
	offset := 12

	for i := 0; i < 32; i++ {
		start := offset + i*2
		end := start + 2
		charges[i] = binary.LittleEndian.Uint16(data[start:end])
	}

	evt := EventData{
		LostBuffer: eventLostBuffer,
		LostFPGA:   eventLostFPGA,
		T0:         t0,
		T1:         t1,
		Charges:    charges,
	}
	return &evt
}

// https://www.geeksforgeeks.org/gray-to-binary-and-binary-to-gray-conversion/
func binToGray(n uint32) uint32 {
	return n ^ (n >> 1)
}

func grayToBin(n uint32) uint32 {
	var res uint32 = n
	for n > 0 {
		n >>= 1
		res ^= n
	}
	return res
}
