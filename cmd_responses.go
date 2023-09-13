package main

import (
	"log"
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

type CitirocProbeRegister struct {
	fastShaperOuput    [32]bool
	SlowShaperLGOutput [32]bool
	peakSensingLG      [32]bool
	SlowShaperHGOutput [32]bool
	peakSensingHG      [32]bool
	preampHGOutput     [32]bool
	preampLGOutput     [32]bool
	inputDAC           [32]bool
}

type CitirocSlowControlRegister struct {
	dac_t                  [32]uint8
	dac                    [32]uint8
	enableDiscriminator    bool
	discriminator          bool
	RS_or_discriminator    bool
	enableDiscriminatorT   bool
	discriminatorT         bool
	enable_4b_dac          bool
	dac4b                  bool
	enable_4b_dac_t        bool
	dac4b_t                bool
	discriminatorMask      [32]bool
	HG_trackHold           bool
	enable_HG_trackHold    bool
	LH_trackHold           bool
	enable_LG_trackHold    bool
	scaBias                bool
	HG_peakDetector        bool
	enable_HG_peakDetector bool
	LG_peakDetector        bool
	enable_LG_peakDetector bool
	sel_SCA_Pdet_HG        bool
	sel_SCA_Pdet_LG        bool
	bypass_PSC             bool
	sel_trig_ext_PSC       bool
	fastShaperFollower     bool
	enableFastShaper       bool
	fastShaper             bool
	LG_slowShaper          bool
	enable_LG_slowShaper   bool
	timeConstant_LG_shaper uint8
	HG_slowShaper          bool
	enable_HG_slowShaper   bool
	timeConstant_HG_shaper uint8
	LG_preamp_bias         bool
	HG_preamp              bool
	enable_HG_praemp       bool
	LG_preamp              bool
	enable_LG_preamp       bool
	fastShaperLG           bool
	enable_input_DAC       bool
	dac8_reference         bool
	input_dac              [32]uint8
	input_dac_on           [32]bool
	channel_preamp_HG      [32]uint8
	channel_preamp_LG      [32]uint8
	channel_Ctest_HG       [32]bool
	channel_Ctest_LG       [32]bool
	channel_preamp_disable [32]bool
	temperature            bool
	enableTemperature      bool
	bandgap                bool
	enable_bandgap         bool
	enable_dac1            bool
	dac1                   bool
	enable_dac2            bool
	dac2                   bool
	dac1_code              uint16
	dac2_code              uint16
	enable_HG_OTA          bool
	HG_OTA                 bool
	enable_LG_OTA          bool
	LG_OTA                 bool
	enable_probe_OTA       bool
	probe_OTA              bool
	testb_OTA              bool
	enable_val_evt         bool
	val_evt                bool
	enable_raz_chn         bool
	raz_chn                bool
	enable_out_dig         bool
	enable_or32            bool
	enable_nor32_oc        bool
	trigger_polarity       bool
	enable_bor32t_oc       bool
	enable_32_triggers     bool
}

func decodeFrame(recvChannel chan Frame, data *DaqData) {
	for {
		frame := <-recvChannel
		log.Printf("length %d, %s", len(frame.Payload), frame.Command)
		switch frame.Command {
		case FEB_OK:
			storeDeviceMac(frame, data)
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
