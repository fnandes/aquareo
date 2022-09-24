package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/pedrobfernandes/aquareo/internal/daemon"
	"github.com/spf13/afero"

	"go.etcd.io/bbolt"
)

func main() {
	db, err := bbolt.Open("aquareo.db", 0600, nil)
	if err != nil {
		log.Fatal("Failed to get a database connection: ", err)
	}
	defer db.Close()

	done := make(chan struct{})
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)

	fs := afero.NewOsFs()
	app := daemon.NewDaemon(fs, db)
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
