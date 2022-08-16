package main

import (
	"context"
	"encoding/json"
	"github.com/pedrobfernandes/aquareo/internal/api"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/pedrobfernandes/aquareo/internal/device"
	"github.com/pedrobfernandes/aquareo/internal/sensor"
	"github.com/pedrobfernandes/aquareo/internal/store"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	config, err := loadConfigFile("config.json")
	if err != nil {
		log.Fatal(err)
	}

	fileStorage := store.NewFileStore()

	controller := device.NewRPiController(fileStorage)
	if err := controller.Open(); err != nil {
		log.Fatal(err)
	}
	defer controller.Close()

	server := api.NewServer(config, controller)
	commander := sensor.NewCommander(config, controller)

	done := make(chan struct{})
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	app := aquareo.NewApp(config, controller, server, commander)

	app.Start()

	go func() {
		defer close(done)
		<-signals

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		app.Stop(ctx)
	}()

	<-done
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
