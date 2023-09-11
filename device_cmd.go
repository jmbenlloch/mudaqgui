package main

import (
	"net"
)

func scanDeviceMac(src net.HardwareAddr, dst net.HardwareAddr) (*Frame, error) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x00, 0x00}) // VCXO
	copy(payload[2:], src)                 // MAC
	// fmt.Println(payload)

	return buildFrame(src, dst, FEB_SET_RECV, payload)
}

func scanDevices(src net.HardwareAddr, sendChannel chan *Frame) {
	for i := 0; i <= 255; i++ {
		dst := net.HardwareAddr{0x00, 0x60, 0x37, 0x12, 0x34, byte(i)}
		frame, _ := scanDeviceMac(src, dst)
		sendChannel <- frame
	}
}
