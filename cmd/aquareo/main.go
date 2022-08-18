package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pedrobfernandes/aquareo/internal/api"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/pedrobfernandes/aquareo/internal/device"
	"github.com/pedrobfernandes/aquareo/internal/sensor"
	"github.com/pedrobfernandes/aquareo/internal/store"
	"go.etcd.io/bbolt"
)

func main() {
	config, err := loadConfigFile("config.json")
	if err != nil {
		log.Fatal("Failed to load the config.json file: ", err)
	}

	db, err := bbolt.Open("aquareo.db", 0600, nil)
	if err != nil {
		log.Fatal("Failed to get a database connection: ", err)
	}
	defer db.Close()

	boltStore := store.NewBoltDbStore(db)
	controller := device.NewRPiController(boltStore)

	if err := controller.Init(config); err != nil {
		log.Fatal("Failed to start controller: ", err)
	}
	defer controller.Close()

	if err := db.Update(func(tx *bbolt.Tx) error {
		for _, s := range controller.Sensors() {
			if _, err := tx.CreateBucketIfNotExists([]byte(s.Id())); err != nil {
				return fmt.Errorf("Failed to create bucket: %w", err)
			}
		}
		return nil
	}); err != nil {
		log.Fatal(err)
	}

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
