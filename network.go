package main

import (
	"fmt"
	"log"
	"net"

	"github.com/mdlayher/packet"
)

func getNetworkInterfacesNames() []string {
	interface_names := make([]string, 0)
	interfaces, _ := net.Interfaces()

	for _, iface := range interfaces {
		interface_names = append(interface_names, iface.Name)
	}
	return interface_names
}

func getNetworkInterface(ifname string) *net.Interface {
	iface, err := net.InterfaceByName(ifname)
	if err != nil {
		log.Fatalf("failed to open interface: %v", err)
	}
	return iface
}

func createSocket(iface *net.Interface) *packet.Conn {
	// Open a raw socket using same EtherType as our frame.
	conn, err := packet.Listen(iface, packet.Raw, DaqEtherType, nil)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("socket type: %T\n", conn)

	return conn
}

func sendFrameViaSocket(sendChannel chan *Frame, connection *packet.Conn) {
	// the Ethernet frame format.
	for {
		frame := <-sendChannel

		b, err := frame.MarshalBinary()
		if err != nil {
			log.Fatalf("failed to marshal frame: %v", err)
		}

		addr := &packet.Addr{HardwareAddr: frame.Destination}
		if _, err := connection.WriteTo(b, addr); err != nil {
			log.Fatalf("failed to write frame: %v", err)
		}
	}
}

// receiveMessages continuously receives messages over a connection. The messages
// may be up to the interface's MTU in size.
func receiveMessages(recvChannel chan Frame, c net.PacketConn, mtu int) {
	log.Println("receive messages go routine")
	var f Frame
	b := make([]byte, mtu)

	// Keep receiving messages forever.
	for {
		n, addr, err := c.ReadFrom(b)
		_ = addr
		if err != nil {
			log.Fatalf("failed to receive message: %v", err)
		}

		// Unpack Ethernet II frame into Go representation.
		if err := (&f).UnmarshalBinary(b[:n]); err != nil {
			log.Fatalf("failed to unmarshal ethernet frame: %v", err)
		}

		recvChannel <- f
	}
}
