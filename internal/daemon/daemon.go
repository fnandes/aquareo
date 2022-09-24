package daemon

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/pedrobfernandes/aquareo/internal/api"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/pedrobfernandes/aquareo/internal/device"
	"github.com/pedrobfernandes/aquareo/internal/sensor"
	"github.com/pedrobfernandes/aquareo/internal/store"
	"github.com/spf13/afero"
	"go.etcd.io/bbolt"
)

type daemon struct {
	fs afero.Fs
	db *bbolt.DB
	ws aquareo.WebServer
	sc aquareo.SensorCommander
}

func NewDaemon(fs afero.Fs, db *bbolt.DB) *daemon {
	return &daemon{
		fs: fs,
		db: db,
	}
}

func (a *daemon) Stop(ctx context.Context) {
	if a.ws != nil {
		a.ws.Stop(ctx)
	}
	if a.sc != nil {
		a.sc.Stop(ctx)
	}
}

func (a *daemon) Start() error {
	s := store.NewBoltDbStore(a.db)
	ctrl := device.NewRPiController(a.fs, s)

	cfg, err := a.loadConfigFile("config.json")
	if err != nil {
		return fmt.Errorf("daemon: Failed to load config file")
	}

	if err := ctrl.Init(cfg); err != nil {
		if !errors.Is(err, device.ErrTempSensorNotFound) {
			return fmt.Errorf("daemon: Failed to start controller: %w", err)
		}
	}

	if err := createBuckets(a.db, ctrl); err != nil {
		return fmt.Errorf("daemon: Failed to create buckets: %w", err)
	}

	a.ws = api.NewServer(cfg, ctrl)
	a.sc = sensor.NewCommander(cfg, ctrl)

	go func() {
		a.ws.Start(":8082")
	}()

	go func() {
		a.sc.Start()
	}()
	return nil
}

func createBuckets(db *bbolt.DB, c aquareo.Controller) error {
	return db.Update(func(tx *bbolt.Tx) error {
		for _, s := range c.Sensors() {
			if _, err := tx.CreateBucketIfNotExists([]byte(s.Id())); err != nil {
				return err
			}
		}
		return nil
	})
}

func (a *daemon) loadConfigFile(filename string) (aquareo.Config, error) {
	var config aquareo.Config

	content, err := afero.ReadFile(a.fs, filename)
	if err != nil {
		return config, err
	}

	if err := json.Unmarshal(content, &config); err != nil {
		return config, err
	}

	return config, nil
}
