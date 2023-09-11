package main

import (
	"encoding/binary"
	"io"
	"net"
)

type EtherType uint16

type Frame struct {
	Destination net.HardwareAddr
	Source      net.HardwareAddr
	EtherType   EtherType
	Command     Cmd

	// Payload is a variable length data payload encapsulated by this Frame.
	Payload []byte
}

const (
	// minPayload is the minimum payload size for an Ethernet frame, assuming
	// that no 802.1Q VLAN tags are present.
	minPayload = 46
)

// read reads data from a Frame into b.  read is used to marshal a Frame
// into binary form, but does not allocate on its own.
func (f *Frame) read(b []byte) (int, error) {
	copy(b[0:6], f.Destination)
	copy(b[6:12], f.Source)

	n := 12

	// Marshal EtherType, Command and copy payload into output bytes.
	binary.BigEndian.PutUint16(b[n:n+2], uint16(f.EtherType))
	binary.LittleEndian.PutUint16(b[n+2:n+4], uint16(f.Command))
	copy(b[n+4:], f.Payload)

	return len(b), nil
}

// length calculates the number of bytes required to store a Frame.
func (f *Frame) length() int {
	// If payload is less than the required minimum length, we zero-pad up to
	// the required minimum length
	pl := len(f.Payload)
	if pl < minPayload {
		pl = minPayload
	}

	// 6 bytes: destination hardware address
	// 6 bytes: source hardware address
	// 2 bytes: EtherType
	// N bytes: payload length (maybe padded)
	return 6 + 6 + 2 + pl
}

// MarshalBinary allocates a byte slice and marshals a Frame into binary form.
func (f *Frame) MarshalBinary() ([]byte, error) {
	b := make([]byte, f.length())
	_, err := f.read(b)
	return b, err
}

// UnmarshalBinary unmarshals a byte slice into a Frame.
func (f *Frame) UnmarshalBinary(b []byte) error {
	// Verify that both hardware addresses and a single EtherType are present
	if len(b) < 14 {
		return io.ErrUnexpectedEOF
	}

	// Track offset in packet for reading data
	n := 14

	// Allocate single byte slice to store destination and source hardware
	// addresses, and payload
	bb := make([]byte, 6+6+len(b[n:]))
	copy(bb[0:6], b[0:6])
	f.Destination = bb[0:6]
	copy(bb[6:12], b[6:12])
	f.Source = bb[6:12]

	f.EtherType = EtherType(binary.BigEndian.Uint16(b[12:14]))
	f.Command = Cmd(binary.LittleEndian.Uint16(b[14:16]))

	// There used to be a minimum payload length restriction here, but as
	// long as two hardware addresses and an EtherType are present, it
	// doesn't really matter what is contained in the payload.  We will
	// follow the "robustness principle".
	copy(bb[12:], b[16:])
	f.Payload = bb[12:]

	return nil
}

func buildFrame(src net.HardwareAddr, dst net.HardwareAddr, command Cmd, payload []byte) (*Frame, error) {
	f := Frame{
		Source:      src,
		Destination: dst,
		EtherType:   DaqEtherType,
		Command:     command,
		Payload:     payload,
	}
	return &f, nil
}
