package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
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
