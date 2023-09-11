package main

import (
	"fmt"
	"log"
	"net"
)

type DaqData struct {
	devices []*net.HardwareAddr
}

func decodeFrame(recvChannel chan *Frame, data *DaqData) {
	for {
		frame := <-recvChannel
		switch frame.Command {
		case FEB_OK:
			storeDeviceMac(frame, data)
		default:
			log.Fatalf("Unkown response command")
		}
	}
}

func storeDeviceMac(frame *Frame, data *DaqData) {
	log.Printf("[%s] %s", frame.Source.String(), string(frame.Payload))
	log.Printf("[%s] %x", frame.Source.String(), frame.Payload)
	log.Printf("[%s] %x", frame.Source.String(), string(frame.EtherType))
	log.Printf("[%s] %x", frame.Source.String(), string(frame.Command))

	data.devices = append(data.devices, &frame.Source)
	fmt.Println(frame.Source[5])
	fmt.Println(data.devices)
}
