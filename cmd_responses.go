package main

import (
	"fmt"
	"log"
	"math"
	"net"

	"encoding/binary"
)

type DaqData struct {
	devices []*net.HardwareAddr
	t0      []uint32
	t1      []uint32
}

type EventData struct {
	//	eventT0    bool
	//	eventT1    bool
	t0         uint32
	t1         uint32
	lostBuffer uint16
	lostFPGA   uint16
	charges    [32]uint16
}

func decodeFrame(recvChannel chan Frame, data *DaqData) {
	for {
		frame := <-recvChannel
		log.Printf("length %d, %s", len(frame.Payload), frame.Command)
		switch frame.Command {
		case FEB_OK:
			storeDeviceMac(frame, data)
			decodeRate(frame, data)
		case FEB_DATA_CDR:
			log.Println("data cdr")
			decodeData(frame, data)
		case FEB_EOF_CDR:
			log.Println("End of data")
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

func storeDeviceMac(frame Frame, data *DaqData) {
	//	log.Printf("[%s] %s", frame.Source.String(), string(frame.Payload))
	//	log.Printf("[%s] %x", frame.Source.String(), frame.Payload)
	//	log.Printf("[%s] %x", frame.Source.String(), string(frame.EtherType))
	//	log.Printf("[%s] %x", frame.Source.String(), string(frame.Command))

	data.devices = append(data.devices, &frame.Source)
	//fmt.Println(frame.Source[5])
	//fmt.Println(data.devices)
}

func decodeData(frame Frame, data *DaqData) {
	//log.Printf("[%s] %s", frame.Source.String(), string(frame.Payload))
	log.Printf("Payload length [%s] %d", frame.Source.String(), len(frame.Payload))

	// Position 0:2 00 00, unused register value

	data_start := 2
	packet_size := 76

	for data_start < len(frame.Payload)-2 {
		log.Printf("reading: %d - %d Len: %d", data_start, data_start+packet_size, len(frame.Payload))
		evt := decodeEvent(frame.Payload[data_start : data_start+packet_size])
		data_start += packet_size

		log.Printf("[Event lost buffer] %d", evt.lostBuffer)
		log.Printf("[Event lost fgpa] %d", evt.lostFPGA)
		log.Printf("[t0] %d", evt.t0)
		log.Printf("[t1] %d", evt.t1)

		for i := 0; i < 32; i++ {
			log.Printf("charge[%d]: %d", i, evt.charges[i])
		}
	}

	//s := hex.EncodeToString(frame.Payload)
	//fmt.Println(s)
	//	log.Printf("[%s] %x", frame.Source.String(), frame.Payload)
	//	log.Printf("[%s] %x", frame.Source.String(), string(frame.EtherType))
	// log.Printf("[%s] %x", frame.Source.String(), string(frame.Command))
}

func decodeRate(frame Frame, data *DaqData) {
	bits := binary.LittleEndian.Uint32(frame.Payload[2:6])
	rate := math.Float32frombits(bits) // in Hz
	fmt.Printf("rate: %v (0x%08x)\n", rate, bits)
}

func decodeEvent(data []byte) *EventData {
	var eventLostFPGA uint16 = binary.LittleEndian.Uint16(data[0:2])
	var eventLostBuffer uint16 = binary.LittleEndian.Uint16(data[2:4])

	var t0 uint32 = binary.LittleEndian.Uint32(data[4:8])
	var t1 uint32 = binary.LittleEndian.Uint32(data[8:12])

	var t0LSB uint32 = (t0 & 0x00000003)
	var t1LSB uint32 = (t1 & 0x00000003)

	t0 = (t0 & 0x3FFFFFF) >> 2
	t1 = (t1 & 0x3FFFFFF) >> 2

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
		lostBuffer: eventLostBuffer,
		lostFPGA:   eventLostFPGA,
		t0:         t0,
		t1:         t1,
		charges:    charges,
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
