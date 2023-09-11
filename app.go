package main

import (
	"context"
	"fmt"
	"net"

	"github.com/mdlayher/packet"
)

// App struct
type App struct {
	ctx              context.Context
	sendFrameChannel chan *Frame
	recvFrameChannel chan *Frame
	data             DaqData
	connection       *packet.Conn
	iface            *net.Interface
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.data = DaqData{
		devices: make([]*net.HardwareAddr, 0, 256),
	}
	a.sendFrameChannel = make(chan *Frame, 2000)
	a.recvFrameChannel = make(chan *Frame, 2000)

	a.iface = getNetworkInterface("enp5s0")
	a.connection = createSocket(a.iface)
}

func (a *App) onshutdown(ctx context.Context) {
	a.connection.Close()
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	go sendFrameViaSocket(a.sendFrameChannel, a.connection)
	go receiveMessages(a.recvFrameChannel, a.connection, a.iface.MTU)
	go decodeFrame(a.recvFrameChannel, &a.data)
	scanDevices(a.iface.HardwareAddr, a.sendFrameChannel)
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// Greet returns a greeting for the given name
func (a *App) ScanDevices() string {
	macAddrs := fmt.Sprint(a.data.devices)
	fmt.Println(macAddrs)
	return macAddrs
}
