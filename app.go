package main

import (
	"context"
	"fmt"
	"net"

	"github.com/mdlayher/packet"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"golang.org/x/exp/maps"
)

// App struct
type App struct {
	ctx              context.Context
	sendFrameChannel chan *Frame
	recvFrameChannel chan Frame
	data             DaqData
	connection       *packet.Conn
	iface            *net.Interface
	dataTaking       bool
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	a.dataTaking = false
	a.data = DaqData{
		devices:                  make(map[byte]*net.HardwareAddr),
		slowControlConfiguration: make(map[byte]map[string]any),
		probeConfiguration:       make(map[byte]map[string]any),
		events:                   make(map[byte][]EventData),
		charges:                  make(map[byte]ChargeHistogram),
		chargesRebinned:          make(map[byte]ChargeHistogram),
	}
	a.sendFrameChannel = make(chan *Frame, 2000)
	a.recvFrameChannel = make(chan Frame, 2000)
}

func (a *App) onshutdown(ctx context.Context) {
	a.connection.Close()
}

func (a *App) ScanDevices() {
	scanDevices(a.iface.HardwareAddr, a.sendFrameChannel)
}

func (a *App) StartRun() {
	startRun(a.iface.HardwareAddr, a.sendFrameChannel)
	a.dataTaking = true
	devices := maps.Values(a.data.devices)
	go readAllCards(a.iface.HardwareAddr, devices, a.sendFrameChannel, a)
}

func (a *App) StopRun() {
	stopRun(a.iface.HardwareAddr, a.sendFrameChannel)
	a.dataTaking = false
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

func (a *App) SetVCXO() {
	dst := net.HardwareAddr{0x00, 0x60, 0x37, 0x12, 0x34, 0x45}
	setVCXO(a.iface.HardwareAddr, dst, a.sendFrameChannel)
}

func (a *App) SetDACThr() {
	dst := net.HardwareAddr{0x00, 0x60, 0x37, 0x12, 0x34, 0x45}
	setDACThr(a.iface.HardwareAddr, dst, a.sendFrameChannel)
}

func (a *App) UpdateCardConfig(card int, slowControl map[string]any, probe map[string]any) {
	fmt.Println(card)
	fmt.Println(slowControl)
	fmt.Println(probe)

	cardID := byte(card)

	for key, value := range slowControl {
		a.data.slowControlConfiguration[cardID][key] = value
	}

	for key, value := range probe {
		a.data.probeConfiguration[cardID][key] = value
	}

	src := a.iface.HardwareAddr
	dst := a.data.devices[cardID]
	updateCardConfig(cardID, &a.data, src, *dst, a.sendFrameChannel)
}

func (a *App) UpdateGlobalConfig(slowControl map[byte]map[string]any, probe map[byte]map[string]any) {
	fmt.Println(slowControl)
	fmt.Println(probe)

	cards := maps.Keys(slowControl)
	for _, card := range cards {
		a.UpdateCardConfig(int(card), slowControl[card], probe[card])
	}
}

func (a *App) SelectConfigFile() string {
	options := runtime.SaveDialogOptions{
		DefaultDirectory:           "/home/jmbenlloch/go/myproject",
		DefaultFilename:            "config.yaml",
		Title:                      "Select file to save configuration",
		TreatPackagesAsDirectories: true,
	}
	file, _ := runtime.SaveFileDialog(a.ctx, options)
	return file
}

func (a *App) SaveConfiguration(file string) {
	saveConfigYaml(&a.data, file)
}

func (a *App) LoadConfiguration(file string) {
	readConfigYaml(&a.data, file)
	fmt.Println("file read")
	sendConfigToUI(&a.data, a.ctx)
	fmt.Println("event sent")
}

func (a *App) WriteDataFile() {
	writeData()
}

func (a *App) GetNetworkInterfaces() []string {
	return getNetworkInterfacesNames()
}

func (a *App) StartConnection(iface string) bool {
	a.iface = getNetworkInterface(iface)
	a.connection = createSocket(a.iface)

	// Start go routines
	go sendFrameViaSocket(a.sendFrameChannel, a.connection)
	go receiveMessages(a.recvFrameChannel, a.connection, a.iface.MTU)
	go decodeFrame(a.recvFrameChannel, &a.data, a.ctx)
	return a.connection != nil
}
