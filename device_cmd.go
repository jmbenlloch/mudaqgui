package main

import (
	"encoding/binary"
	"fmt"
	"net"
	"time"

	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2/pkg/runtime"
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
	for i := 0; i <= 255; i++ {
		dst := getMacAddressDevice(byte(i))
		frame, _ := scanDeviceMac(src, dst)
		sendChannel <- frame
		time.Sleep(10 * time.Millisecond)
	}
	time.Sleep(1 * time.Second)
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

func readAllCards(src net.HardwareAddr, devices []*net.HardwareAddr, sendChannel chan *Frame, a *App) {
	counter := 0
	for a.dataTaking {
		fmt.Println(a.dataTaking)
		fmt.Printf("len channel: %v\n", len(sendChannel))
		time.Sleep(50 * time.Millisecond)
		for _, dst := range devices {
			getRate(src, *dst, sendChannel)
			readData(src, *dst, sendChannel)

			if counter%20 == 0 {
				runtime.EventsEmit(a.ctx, "rate", a.data.rates)

				runtime.EventsEmit(a.ctx, "events", a.data.events)
				runtime.EventsEmit(a.ctx, "charges", a.data.charges)
				runtime.EventsEmit(a.ctx, "chargesRebin", a.data.chargesRebinned)
			}
			counter = counter + 1
		}
	}
}

func sendProbeConfiguration(configuration map[string]any, src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+(256/8)) // register + mac address
	fmt.Println("probe register")
	fmt.Println(configuration)
	fmt.Println("finished")
	configurationBytes := configurationToByteArray(256, configuration, citirocProbeBitPosition)

	for i, value := range configurationBytes {
		if (i%16 == 0) && (i > 0) {
			fmt.Printf("\n")
		}
		fmt.Printf("%02x ", value)
	}
	copy(payload[0:2], []byte{0x00, 0x00}) //
	copy(payload[2:], configurationBytes)  //
	frame, _ := buildFrame(src, dst, FEB_WR_PMR, payload)
	sendChannel <- frame
}

func sendSlowControlConfiguration(configuration map[string]any, src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+(1144/8))
	configurationBytes := configurationToByteArray(1144, configuration, citirocSlowControlBitPosition)

	for i, value := range configurationBytes {
		if (i%16 == 0) && (i > 0) {
			fmt.Printf("\n")
		}
		fmt.Printf("%02x ", value)
	}
	copy(payload[0:2], []byte{0x00, 0x00}) //
	copy(payload[2:], configurationBytes)  //
	frame, _ := buildFrame(src, dst, FEB_WR_SCR, payload)
	sendChannel <- frame
}

func sendFPGAFil(slowControlConfiguration map[string]any, src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+9)
	mask, ok := slowControlConfiguration["discriminatorMask"].([32]int)
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

func updateCardConfig(card byte, data *DaqData, src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	sendSlowControlConfiguration(data.slowControlConfiguration[card], src, dst, sendChannel)
	sendProbeConfiguration(data.probeConfiguration[card], src, dst, sendChannel)
	sendFPGAFil(data.slowControlConfiguration[card], src, dst, sendChannel)
}

func setVCXO(vcxoValue uint16, src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	payload := make([]byte, 2+6) // register + mac address
	bits := uint16ToByteArray(vcxoValue)
	copy(payload[0:2], bits) // VCXO
	copy(payload[2:], src)   // MAC
	frame, _ := buildFrame(src, dst, FEB_SET_RECV, payload)
	sendChannel <- frame
}

func setDAC1Thr(card byte, data *DaqData, dacValue uint16, src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	data.slowControlConfiguration[card]["dac1_code"] = dacValue
	sendSlowControlConfiguration(data.slowControlConfiguration[card], src, dst, sendChannel)
}

func setDAC2Thr(card byte, data *DaqData, dacValue uint16, src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	data.slowControlConfiguration[card]["dac2_code"] = dacValue
	sendSlowControlConfiguration(data.slowControlConfiguration[card], src, dst, sendChannel)
}

func setDACThr(card byte, data *DaqData, dacValue uint16, src net.HardwareAddr, dst net.HardwareAddr, sendChannel chan *Frame) {
	setDAC1Thr(card, data, dacValue, src, dst, sendChannel)
	setDAC2Thr(card, data, dacValue, src, dst, sendChannel)
}

///////////
// utils //
///////////

func convertToBinary(values [32]int) uint32 {
	var result uint32 = 0
	for i := 0; i < len(values); i++ {
		bit := uint32(values[i] & 0x00000001)
		result = result | (bit << (31 - i))
	}
	return result
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
