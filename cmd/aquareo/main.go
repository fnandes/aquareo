package main

import (
	"encoding/json"
	"github.com/pedrobfernandes/aquareo/internal/api"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/pedrobfernandes/aquareo/internal/device"
	"github.com/pedrobfernandes/aquareo/internal/io"
	"github.com/pedrobfernandes/aquareo/internal/sensor"
	"io/ioutil"
	"log"
	"sync"
)

func main() {
	config, err := loadConfigFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	fileStorage := io.NewFileStorage()

	controller := device.NewRPiController(fileStorage)
	if err := controller.Open(); err != nil {
		log.Fatal(err)
	}
	defer controller.Close()

	server := api.NewServer(config, controller)
	commander := sensor.NewCommander(config)

	app := aquareo.NewApp(config, controller, server, commander)

	var wg sync.WaitGroup
	defer wg.Wait()

	app.Start(&wg)
}

func loadConfigFile(filename string) (aquareo.Config, error) {
	var config aquareo.Config

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(content, &config); err != nil {
		return config, err
	}

	return config, nil
}
