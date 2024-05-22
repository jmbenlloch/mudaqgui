package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"log"
	"net"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

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
		t0s:                      make(map[byte][]uint32),
		t1s:                      make(map[byte][]uint32),
		lostBuffer:               make(map[byte]uint32),
		lostFGPA:                 make(map[byte]uint32),
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

func getOutputFilename() string {
	fmt.Println("MUON DATA:", os.Getenv("MUONDATA"))
	basepath := os.Getenv("MUONDATA")
	pattern := filepath.Join(basepath, "muons_run_*.h5")
	files, _ := filepath.Glob(pattern)

	fmt.Println(files)

	var latest_run int64 = 1

	if len(files) > 0 {
		for i := 0; i < len(files); i++ {
			basename := strings.Split(files[i], ".")[0]
			fnumberStr := strings.Split(basename, "_")[2]
			fnumber, _ := strconv.ParseInt(fnumberStr, 10, 32)
			fmt.Println("file: ", fnumber)
			if fnumber >= latest_run {
				latest_run = fnumber + 1
				fmt.Println("latest run: ", latest_run)
			}
		}
	}

	filename := fmt.Sprintf("%s/muons_run_%d.h5", basepath, latest_run)
	return filename
}

func createOutputFile(writerData *WriterData) {
	fname := getOutputFilename()
	h5file := openFile(fname)
	writerData.file = h5file
	dataset := createTable(writerData.file)
	chargesArray := createChargesArray(writerData.file)
	writerData.data = dataset
	writerData.charges = chargesArray
}

func closeOutputFile(writerData *WriterData) {
	fmt.Println("closing output file")
	writerData.file.Close()
	writerData.data.Close()
	writerData.charges.Close()
}

func (a *App) StartRun() {
	startRun(a.iface.HardwareAddr, a.sendFrameChannel)
	a.dataTaking = true
	runtime.EventsEmit(a.ctx, "dataTaking", a.dataTaking)
	devices := maps.Values(a.data.devices)

	for _, device := range devices {
		card := (*device)[5]
		initialize_charge_histograms(card, &a.data)
		a.data.t0s[card] = make([]uint32, 0)
		a.data.t1s[card] = make([]uint32, 0)
		a.data.lostBuffer[card] = 0
		a.data.lostFGPA[card] = 0
	}

	createOutputFile(&a.writerData)
	go readAllCards(a.iface.HardwareAddr, devices, a.sendFrameChannel, a)
}

func (a *App) StopRun() {
	stopRun(a.iface.HardwareAddr, a.sendFrameChannel)
	// Write remaining events
	writeData(a.writerData.data, &a.data.events)
	writeCharges(a.writerData.charges, &a.data.events)
	a.data.events = make([]EventData, 0, 10000)

	// Close file and stop data taking
	if a.dataTaking {
		// Close the file only once
		closeOutputFile(&a.writerData)
	}
	a.dataTaking = false
	runtime.EventsEmit(a.ctx, "dataTaking", a.dataTaking)
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
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	options := runtime.SaveDialogOptions{
		DefaultDirectory:           directory,
		DefaultFilename:            "config.yaml",
		Title:                      "Select file to save configuration",
		TreatPackagesAsDirectories: true,
	}
	file, _ := runtime.SaveFileDialog(a.ctx, options)
	return file
}

func (a *App) SelectCalibrationFile() string {
	directory, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
	}
	options := runtime.OpenDialogOptions{
		DefaultDirectory:           directory,
		DefaultFilename:            "config.yaml",
		Title:                      "Select file to save configuration",
		TreatPackagesAsDirectories: true,
	}
	file, _ := runtime.OpenFileDialog(a.ctx, options)
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

type CalibrationConfig struct {
	Duration int
	Events   int
	Bias     int
	Gain     int
	Dac      int
}

type CalibrationLog struct {
	Timestamp     int
	Configuration CalibrationConfig
}

func (a *App) LoadCalibrationFile(filename string) {
	fmt.Println(filename)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal("Error while reading the file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()

	// Checks for the error
	if err != nil {
		fmt.Println("Error reading records")
	}

	// Loop to iterate through
	// and print each of the string slice
	for index, record := range records {
		fmt.Println(record)
		if index > 0 {

			duration, err := strconv.Atoi(record[0])
			if err != nil {
				fmt.Println("error reading calibration file")
			}
			events, err := strconv.Atoi(record[1])
			if err != nil {
				fmt.Println("error reading calibration file")
			}
			bias, err := strconv.Atoi(record[2])
			if err != nil {
				fmt.Println("error reading calibration file")
			}
			gain, err := strconv.Atoi(record[3])
			if err != nil {
				fmt.Println("error reading calibration file")
			}
			dac, err := strconv.Atoi(record[4])
			if err != nil {
				fmt.Println("error reading calibration file")
			}
			config := CalibrationConfig{
				Duration: duration,
				Events:   events,
				Bias:     bias,
				Gain:     gain,
				Dac:      dac,
			}
			fmt.Println(config)

			for card, configuration := range a.data.slowControlConfiguration {
				configuration["dac1_code"] = config.Dac
				configuration["dac2_code"] = config.Dac

				gains := make([]int, 32)
				for i := range gains {
					gains[i] = config.Gain
				}
				biases := make([]int, 32)
				for i := range biases {
					biases[i] = config.Bias
				}
				configuration["channel_preamp_HG"] = gains
				configuration["input_dac"] = biases

				fmt.Println(card, configuration)
			}

			runtime.EventsEmit(a.ctx, "configSlowControl", a.data.slowControlConfiguration)
			a.UpdateGlobalConfig(a.data.slowControlConfiguration, a.data.probeConfiguration)
			runtime.EventsEmit(a.ctx, "calibration", CalibrationLog{Timestamp: int(time.Now().Unix()), Configuration: config})
			a.StartRun()
			time.Sleep(time.Duration(config.Duration) * time.Second)
			a.StopRun()
		}
	}
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
