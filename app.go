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
	writerData       WriterData
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
		rates:                    make(map[byte]float32),
		cards:                    make(map[byte]bool),
		events:                   make([]EventData, 0, 10000),
		charges:                  make(map[byte]ChargeHistogram),
		chargesRebinned:          make(map[byte]ChargeHistogram),
	}
	a.writerData = WriterData{}
	a.sendFrameChannel = make(chan *Frame, 2000)
	a.recvFrameChannel = make(chan Frame, 2000)
}

func (a *App) onshutdown(ctx context.Context) {
	a.connection.Close()
}

func (a *App) ScanDevices() {
	scanDevices(a.iface.HardwareAddr, a.sendFrameChannel)
}

func createOutputFile(writerData *WriterData) {
	fname := "/home/jmbenlloch/go/myproject/testfile.h5"
	h5file := openFile(fname)
	writerData.file = h5file
	dataset := createTable(writerData.file)
	chargesArray := createChargesArray(writerData.file)
	writerData.data = dataset
	writerData.charges = chargesArray
}

func closeOutputFile(writerData *WriterData) {
	writerData.file.Close()
	writerData.data.Close()
	writerData.charges.Close()
}

func (a *App) StartRun() {
	startRun(a.iface.HardwareAddr, a.sendFrameChannel)
	a.dataTaking = true
	devices := maps.Values(a.data.devices)
	createOutputFile(&a.writerData)
	go readAllCards(a.iface.HardwareAddr, devices, a.sendFrameChannel, a)
}

func (a *App) StopRun() {
	stopRun(a.iface.HardwareAddr, a.sendFrameChannel)
	closeOutputFile(&a.writerData)
	a.dataTaking = false
}

func (a *App) HVOn(cardID byte) {
	dst := getMacAddressDevice(cardID)
	hvOn(a.iface.HardwareAddr, dst, a.sendFrameChannel)
}

func (a *App) HVOff(cardID byte) {
	dst := getMacAddressDevice(cardID)
	hvOff(a.iface.HardwareAddr, dst, a.sendFrameChannel)
}

func (a *App) SetVCXO(cardID byte, vcxoValue uint16) {
	dst := getMacAddressDevice(cardID)
	fmt.Println("set vcxo")
	setVCXO(vcxoValue, a.iface.HardwareAddr, dst, a.sendFrameChannel)
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

func (a *App) GetNetworkInterfaces() []string {
	return getNetworkInterfacesNames()
}

func (a *App) StartConnection(iface string) bool {
	a.iface = getNetworkInterface(iface)
	a.connection = createSocket(a.iface)

	// Start go routines
	go sendFrameViaSocket(a.sendFrameChannel, a.connection)
	go receiveMessages(a.recvFrameChannel, a.connection, a.iface.MTU)
	go decodeFrame(a.recvFrameChannel, &a.data, &a.writerData, a.ctx)
	return a.connection != nil
}

func (a *App) SetDACThreshold(card byte, dacValue uint16) {
	dst := getMacAddressDevice(card)
	setDACThr(card, &a.data, dacValue, a.iface.HardwareAddr, dst, a.sendFrameChannel)
}
