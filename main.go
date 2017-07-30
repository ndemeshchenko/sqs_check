package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	sqscheck "./lib"
)

func readConfig() (*sqscheck.Config, error) {
	file, err := ioutil.ReadFile("./config.json")
	if err != nil {
		log.Fatalf("File error: %v\n", err)
	}
	var config *sqscheck.Config
	json.Unmarshal(file, &config)
	return config, err
}

func main() {
	config, err := readConfig()
	if err != nil {
		log.Fatal("Unable to read config: config.json")
	}
	sqscheck.Run(config)

}
