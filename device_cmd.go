package main

import (
	"fmt"
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

func getRate(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x00, 0x00}) //
	copy(payload[2:], src)                 // MAC
	frame, _ := buildFrame(src, dst, FEB_GET_RATE, payload)
	sendChannel <- frame
}

func startAcquisition(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x01, 0x00}) // VCXO
	copy(payload[2:], src)                 // MAC
	frame, _ := buildFrame(src, dst, FEB_GEN_INIT, payload)
	sendChannel <- frame

	payload = make([]byte, 2+6)            // register + mac address
	copy(payload[0:2], []byte{0x02, 0x00}) // VCXO
	copy(payload[2:], src)                 // MAC
	frame, _ = buildFrame(src, dst, FEB_GEN_INIT, payload)

	sendChannel <- frame
}

func hvOn(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x02, 0x02}) // VCXO
	copy(payload[2:], src)                 // MAC
	frame, _ := buildFrame(src, dst, FEB_GEN_HVON, payload)
	sendChannel <- frame
}

func hvOff(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x02, 0x02}) // VCXO
	copy(payload[2:], src)                 // MAC
	frame, _ := buildFrame(src, dst, FEB_GEN_HVOFF, payload)
	sendChannel <- frame
}

func readData(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)                 // register + mac address
	copy(payload[0:3], []byte{0x00, 0x00, 0x00}) // VCXO
	copy(payload[3:], src)                       // MAC
	frame, _ := buildFrame(src, dst, FEB_RD_CDR, payload)
	sendChannel <- frame
}

func sendProbeConfiguration(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+(256/8)) // register + mac address
	configuration := configurationToByteArray(256, probeRegister, citirocProbeBitPosition)

	for i, value := range configuration {
		if (i%16 == 0) && (i > 0) {
			fmt.Printf("\n")
		}
		fmt.Printf("%02x ", value)
	}
	copy(payload[0:2], []byte{0x00, 0x00}) //
	copy(payload[2:], configuration)       //
	frame, _ := buildFrame(src, dst, FEB_WR_PMR, payload)
	sendChannel <- frame
}

func sendSlowControlConfiguration(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+(1144/8))
	configuration := configurationToByteArray(1144, slowControlRegister, citirocSlowControlBitPosition)

	for i, value := range configuration {
		if (i%16 == 0) && (i > 0) {
			fmt.Printf("\n")
		}
		fmt.Printf("%02x ", value)
	}
	copy(payload[0:2], []byte{0x00, 0x00}) //
	copy(payload[2:], configuration)       //
	frame, _ := buildFrame(src, dst, FEB_WR_SCR, payload)
	sendChannel <- frame
}

func updateConfig(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	sendSlowControlConfiguration(src, dst, sendChannel)
	sendProbeConfiguration(src, dst, sendChannel)
}
