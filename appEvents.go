package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/wailsapp/wails/v2/pkg/runtime"
	"gopkg.in/yaml.v3"
)

func sendConfigToUI(data *DaqData, ctx context.Context) {
	fmt.Println("event emitted1")
	slowControl := make(map[byte]map[string]any)
	probe := make(map[byte]map[string]any)

	for card, config := range data.slowControlConfiguration {
		partialConfig := make(map[string]any)
		partialConfig["channel_preamp_HG"] = config["channel_preamp_HG"]
		partialConfig["input_dac"] = config["input_dac"]
		partialConfig["channel_preamp_disable"] = config["channel_preamp_disable"]
		partialConfig["discriminatorMask"] = config["discriminatorMask"]
		partialConfig["enable_or32"] = config["enable_or32"]
		partialConfig["dac1_code"] = config["dac1_code"]
		partialConfig["dac2_code"] = config["dac2_code"]

		slowControl[card] = partialConfig
	}

	for card, config := range data.probeConfiguration {
		partialConfig := make(map[string]any)
		partialConfig["peakSensingHG"] = config["peakSensingHG"]

		probe[card] = partialConfig
	}

	runtime.EventsEmit(ctx, "configSlowControl", slowControl)
	runtime.EventsEmit(ctx, "configProbe", probe)
	fmt.Println("event emitted")
}

func readConfigFromUI(data *DaqData, ctx context.Context) {

}

func saveConfigYaml(data *DaqData, file string) {
	configToSave := make(map[byte]map[string]any)

	for card, values := range data.probeConfiguration {
		configToSave[card] = make(map[string]any)
		configToSave[card]["probeConfiguration"] = values
	}

	for card, values := range data.slowControlConfiguration {
		configToSave[card]["slowControlConfiguration"] = values
	}

	dataTest, err := yaml.Marshal(configToSave)
	if err != nil {
		log.Fatal(err)
	}

	err2 := os.WriteFile(file, dataTest, 0644)
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println("data written")
	_ = dataTest
}

func readConfigYaml(data *DaqData, file string) {
	yfile, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	dataYaml := make(map[byte]map[string]map[string]any)
	err2 := yaml.Unmarshal(yfile, &dataYaml)
	_ = err2

	for card, values := range dataYaml {
		data.probeConfiguration[card] = values["probeConfiguration"]
		data.slowControlConfiguration[card] = values["slowControlConfiguration"]
	}
}
