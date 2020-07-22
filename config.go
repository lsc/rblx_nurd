package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"time"

	"fmt"
)

type ConfigFile struct {
	VictoriaMetrics Server
	Nomad           []Server
}

type Server struct {
	URL  string
	Port string
}

var (
	nomadAddresses []string
	metricsAddressPointer string
	duration time.Duration
)

func loadConfig(path string) {
	fmt.Println("Begin loadConfig")
	nomadAddresses = []string{}

	data, err := ioutil.ReadFile(path)
	var config ConfigFile
	err = json.Unmarshal(data, &config)
	if err != nil {
		log.Fatal(err)
	}
	
	metricsAddressPointer = config.VictoriaMetrics.URL + ":" + config.VictoriaMetrics.Port

	for _, server := range config.Nomad {
		nomadAddresses = append(nomadAddresses, server.URL+":"+server.Port)
	}

	duration, err = time.ParseDuration("1m")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finish loadConfig")
}
