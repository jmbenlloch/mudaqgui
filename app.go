package main

import (
	"context"
	"log"
	"net"

	"github.com/mdlayher/packet"
)

// App struct
type App struct {
	ctx              context.Context
	sendFrameChannel chan *Frame
	recvFrameChannel chan Frame
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
		devices: make(map[byte]*net.HardwareAddr),
		t0:      make([]uint32, 100000),
		t1:      make([]uint32, 100000),
	}
	a.sendFrameChannel = make(chan *Frame, 2000)
	a.recvFrameChannel = make(chan Frame, 2000)

	a.iface = getNetworkInterface("enp5s0")
	a.connection = createSocket(a.iface)

	// Start go routines
	go sendFrameViaSocket(a.sendFrameChannel, a.connection)
	go receiveMessages(a.recvFrameChannel, a.connection, a.iface.MTU)
	go decodeFrame(a.recvFrameChannel, &a.data)
}

func (a *App) onshutdown(ctx context.Context) {
	a.connection.Close()
}

func (a *App) ScanDevices() {
	scanDevices(a.iface.HardwareAddr, a.sendFrameChannel)
}

func (a *App) DevicesMacs() []string {
	macAddrs := make([]string, 0, 256)
	for _, mac := range a.data.devices {
		macAddrs = append(macAddrs, mac.String())
	}
	return macAddrs
}

func (a *App) GetRate() {
	dst := net.HardwareAddr{0x00, 0x60, 0x37, 0x12, 0x34, 0x45}
	getRate(a.iface.HardwareAddr, dst, a.sendFrameChannel)
}

func (a *App) StartRun() {
	startRun(a.iface.HardwareAddr, a.sendFrameChannel)
}

func (a *App) StopRun() {
	stopRun(a.iface.HardwareAddr, a.sendFrameChannel)
}

func (a *App) HVOn() {
	dst := net.HardwareAddr{0x00, 0x60, 0x37, 0x12, 0x34, 0x45}
	hvOn(a.iface.HardwareAddr, dst, a.sendFrameChannel)
}

func (a *App) HVOff() {
	dst := net.HardwareAddr{0x00, 0x60, 0x37, 0x12, 0x34, 0x45}
	hvOff(a.iface.HardwareAddr, dst, a.sendFrameChannel)
}

func (a *App) ReadData() {
	dst := net.HardwareAddr{0x00, 0x60, 0x37, 0x12, 0x34, 0x45}
	readData(a.iface.HardwareAddr, dst, a.sendFrameChannel)
}

func (a *App) UpdateConfig() {
	dst := net.HardwareAddr{0x00, 0x60, 0x37, 0x12, 0x34, 0x45}
	updateConfig(a.iface.HardwareAddr, dst, a.sendFrameChannel)
}

func (a *App) SetVCXO() {
	dst := net.HardwareAddr{0x00, 0x60, 0x37, 0x12, 0x34, 0x45}
	setVCXO(a.iface.HardwareAddr, dst, a.sendFrameChannel)
}

func (a *App) SetDACThr() {
	dst := net.HardwareAddr{0x00, 0x60, 0x37, 0x12, 0x34, 0x45}
	setDACThr(a.iface.HardwareAddr, dst, a.sendFrameChannel)
}

func (a *App) PrintT0() {
	for _, t0 := range a.data.t0 {
		log.Println(t0)
	}
}
