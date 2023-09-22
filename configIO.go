package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

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
