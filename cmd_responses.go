package main

import (
	"context"
	"log"
	"math"
	"net"

	"encoding/binary"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

type CardRate struct {
	Card byte    `json:"card"`
	Rate float32 `json:"rate"`
}

type DaqData struct {
	devices                  map[byte]*net.HardwareAddr
	slowControlConfiguration map[byte]map[string]any
	probeConfiguration       map[byte]map[string]any
	rates                    map[byte]float32
	cards                    map[byte]bool
	events                   []EventData
	charges                  map[byte]ChargeHistogram
	chargesRebinned          map[byte]ChargeHistogram
	t0s                      map[byte][]uint32
	t1s                      map[byte][]uint32
	lostBuffer               map[byte]uint32
	lostFGPA                 map[byte]uint32
}

type EventData struct {
	card       byte
	eventT0    bool
	eventT1    bool
	overflowT0 bool
	overflowT1 bool
	T0         uint32     `json="t0"`
	T1         uint32     `json="t1"`
	LostBuffer uint16     `json="lostBuffer"`
	LostFPGA   uint16     `json="lostFGPA"`
	Charges    [32]uint16 `json="charges"`
}

type ChargeHistogram struct {
	Charges [32][]int32 `json="charges"`
}

func decodeFrame(recvChannel chan Frame, data *DaqData, writerData *WriterData, ctx context.Context) {
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
			writeEvents(data, writerData)
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

	if _, exists := data.cards[frame.Source[5]]; !exists {
		// Add default config and send to UI
		data.slowControlConfiguration[frame.Source[5]] = createDefaultSlowControlConfiguration()
		data.probeConfiguration[frame.Source[5]] = createDefaultProbeRegisterConfiguration()
		sendConfigToUI(data, ctx)

		// Add card and send to UI
		data.cards[frame.Source[5]] = true
		cards := make([]int, 0, len(data.devices))
		for key := range data.devices {
			cards = append(cards, int(key))
		}
		// For some reason these values are not properly passed to JS
		//runtime.EventsEmit(ctx, "cards", maps.Keys(data.devices))
		runtime.EventsEmit(ctx, "cards", cards)
	}

	// Add new list to events if this card did not exist
	//if _, exists := data.events[frame.Source[5]]; !exists {
	//data.events[frame.Source[5]] = make([]EventData, 0, 10000)
	//}

	// Add new charge histogram if this card did not exist
	if _, exists := data.charges[frame.Source[5]]; !exists {
		initialize_charge_histograms(frame.Source[5], data)
	}

	if _, exists := data.t0s[frame.Source[5]]; !exists {
		data.t0s[frame.Source[5]] = make([]uint32, 0)
		data.t1s[frame.Source[5]] = make([]uint32, 0)
	}
}

func initialize_charge_histograms(card byte, data *DaqData) {
	data.charges[card] = *create_charge_histograms(4096)
	data.chargesRebinned[card] = *create_charge_histograms(128)
}

func create_charge_histograms(nbins int) *ChargeHistogram {
	chargeHistograms := ChargeHistogram{}
	for i := 0; i < 32; i++ {
		chargeHistograms.Charges[i] = make([]int32, nbins)
	}
	return &chargeHistograms
}

func decodeData(frame Frame, data *DaqData, ctx context.Context) {
	//log.Printf("[%s] %s", frame.Source.String(), string(frame.Payload))
	log.Printf("Payload length [%s] %d", frame.Source.String(), len(frame.Payload))

	// Position 0:2 00 00, unused register value

	data_start := 2
	packet_size := 76

	for data_start < len(frame.Payload)-2 {
		//log.Printf("reading: %d - %d Len: %d", data_start, data_start+packet_size, len(frame.Payload))
		evt := decodeEvent(frame.Payload[data_start : data_start+packet_size])
		data_start += packet_size

		//data.events[frame.Source[5]] = append(data.events[frame.Source[5]], *evt)
		evt.card = frame.Source[5]
		data.events = append(data.events, *evt)

		nEvts := len(data.t0s[evt.card])
		start := 0
		if (nEvts - 1000) > 0 {
			start = nEvts - 1000
		}
		data.t0s[evt.card] = append(data.t0s[evt.card][start:], evt.T0)
		data.t1s[evt.card] = append(data.t1s[evt.card][start:], evt.T1)

		data.lostBuffer[evt.card] = data.lostBuffer[evt.card] + uint32(evt.LostBuffer)
		data.lostFGPA[evt.card] = data.lostFGPA[evt.card] + uint32(evt.LostFPGA)

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
			index := evt.Charges[i] / 32
			//fmt.Println(index, evt.Charges[i])
			count = chargesHistograms.Charges[i][index]
			chargesHistograms.Charges[i][index] = count + 1
		}
	}

	// s := hex.EncodeToString(frame.Payload)
	// fmt.Println(s)
	//
	//	log.Printf("[%s] %x", frame.Source.String(), frame.Payload)
	//	log.Printf("[%s] %x", frame.Source.String(), string(frame.EtherType))
	//
	// log.Printf("[%s] %x", frame.Source.String(), string(frame.Command))
}

func writeEvents(data *DaqData, writerData *WriterData) {
	if len(data.events) > 100 {
		writeData(writerData.data, &data.events)
		writeCharges(writerData.charges, &data.events)
		data.events = make([]EventData, 0, 10000)
	}
}

func decodeRate(frame Frame, data *DaqData, ctx context.Context) {
	bits := binary.LittleEndian.Uint32(frame.Payload[2:6])
	rate := math.Float32frombits(bits) // in Hz
	data.rates[frame.Source[5]] = rate
}

func decodeEvent(data []byte) *EventData {
	var eventLostFPGA uint16 = binary.LittleEndian.Uint16(data[0:2])
	var eventLostBuffer uint16 = binary.LittleEndian.Uint16(data[2:4])

	var t0 uint32 = binary.LittleEndian.Uint32(data[4:8])
	var t1 uint32 = binary.LittleEndian.Uint32(data[8:12])

	var eventT0 bool = (t0 & 0x80000000) > 0
	var eventT1 bool = (t1 & 0x80000000) > 0

	var overflowT0 bool = (t0 & 0x40000000) > 0
	var overflowT1 bool = (t1 & 0x40000000) > 0

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
		eventT0:    eventT0,
		eventT1:    eventT1,
		overflowT0: overflowT0,
		overflowT1: overflowT1,
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
