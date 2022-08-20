package main

import (
	"context"
	"encoding/json"
	"github.com/pedrobfernandes/aquareo/internal/daemon"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
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

	done := make(chan struct{})
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	app := daemon.NewDaemon(config, db)
	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

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
