package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/labstack/gommon/log"
)

func getMacAddressDevice(device byte) net.HardwareAddr {
	mac := net.HardwareAddr{0x00, 0x60, 0x37, 0x12, 0x34, device}
	return mac
}

func scanDeviceMac(src net.HardwareAddr, dst net.HardwareAddr) (*Frame, error) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x00, 0x00}) // VCXO
	copy(payload[2:], src)                 // MAC
	return buildFrame(src, dst, FEB_SET_RECV, payload)
}

func scanDevices(src net.HardwareAddr, sendChannel chan *Frame) {
	// Scan several times. For some reason boards do not always answer (timing?)
	for scan := 0; scan <= 2; scan++ {
		for i := 0; i <= 255; i++ {
			dst := getMacAddressDevice(byte(i))
			frame, _ := scanDeviceMac(src, dst)
			sendChannel <- frame
			time.Sleep(10 * time.Millisecond)
		}
		time.Sleep(1 * time.Second)
	}
}

func getRate(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x00, 0x00}) // register
	copy(payload[2:], src)                 // MAC
	frame, _ := buildFrame(src, dst, FEB_GET_RATE, payload)
	sendChannel <- frame
}

func startRun(src net.HardwareAddr, sendChannel chan *Frame) {
	// Send to broadcast
	dst := getMacAddressDevice(0xff)

	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x01, 0x00}) // reset
	copy(payload[2:], src)                 // MAC
	frame, _ := buildFrame(src, dst, FEB_GEN_INIT, payload)
	sendChannel <- frame

	payload = make([]byte, 2+6)            // register + mac address
	copy(payload[0:2], []byte{0x02, 0x00}) // start data acquisition
	copy(payload[2:], src)                 // MAC
	frame, _ = buildFrame(src, dst, FEB_GEN_INIT, payload)

	sendChannel <- frame
}

func stopRun(src net.HardwareAddr, sendChannel chan *Frame) {
	// Send to broadcast
	dst := getMacAddressDevice(0xff)
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x00, 0x00}) // stop data acquisition
	copy(payload[2:], src)                 // MAC
	frame, _ := buildFrame(src, dst, FEB_GEN_INIT, payload)

	sendChannel <- frame
}

func hvOn(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x02, 0x02}) //
	copy(payload[2:], src)                 // MAC
	frame, _ := buildFrame(src, dst, FEB_GEN_HVON, payload)
	sendChannel <- frame
}

func hvOff(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)           // register + mac address
	copy(payload[0:2], []byte{0x02, 0x02}) //
	copy(payload[2:], src)                 // MAC
	frame, _ := buildFrame(src, dst, FEB_GEN_HVOFF, payload)
	sendChannel <- frame
}

func readData(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+6)                 // register + mac address
	copy(payload[0:3], []byte{0x00, 0x00, 0x00}) //
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

func sendFPGAFil(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+9)
	mask, ok := slowControlRegister["discriminatorMask"].([32]int)
	if !ok {
		log.Info("error")
	}
	binaryMask := convertToBinary(mask)
	bits := uint32ToByteArray(binaryMask)
	copy(payload[0:2], []byte{0x00, 0x00}) //
	copy(payload[2:], bits)                //
	frame, _ := buildFrame(src, dst, FEB_WR_FIL, payload)
	sendChannel <- frame
}

func updateConfig(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	sendSlowControlConfiguration(src, dst, sendChannel)
	sendProbeConfiguration(src, dst, sendChannel)
	sendFPGAFil(src, dst, sendChannel)
}

func uint16ToByteArray(value uint16) []byte {
	array := make([]byte, 2)
	binary.LittleEndian.PutUint16(array, value)
	return array
}

func uint32ToByteArray(value uint32) []byte {
	array := make([]byte, 4)
	binary.LittleEndian.PutUint32(array, value)
	return array
}

func setVCXO(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+6) // register + mac address
	var vcxo uint16 = 1023
	bits := uint16ToByteArray(vcxo)
	copy(payload[0:2], bits) // VCXO
	copy(payload[2:], src)   // MAC
	frame, _ := buildFrame(src, dst, FEB_GEN_HVOFF, payload)
	sendChannel <- frame
}

func setDAC1Thr(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	slowControlRegister["dac1_code"] = 768
	sendSlowControlConfiguration(src, dst, sendChannel)
}

func setDAC2Thr(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	slowControlRegister["dac2_code"] = 768
	sendSlowControlConfiguration(src, dst, sendChannel)
}

func setDACThr(src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	setDAC1Thr(src, dst, sendChannel)
	setDAC2Thr(src, dst, sendChannel)
}

func convertToBinary(values [32]int) uint32 {
	var result uint32 = 0
	for i := 0; i < len(values); i++ {
		bit := uint32(values[i] & 0x00000001)
		result = result | (bit << (31 - i))
	}
	return result
}
