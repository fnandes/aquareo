package daemon

import (
	"context"
	"fmt"
	"github.com/pedrobfernandes/aquareo/internal/api"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/pedrobfernandes/aquareo/internal/device"
	"github.com/pedrobfernandes/aquareo/internal/sensor"
	"github.com/pedrobfernandes/aquareo/internal/store"
	"go.etcd.io/bbolt"
)

type daemon struct {
	cfg aquareo.Config
	db  *bbolt.DB
	ws  aquareo.WebServer
	sc  aquareo.SensorCommander
}

func NewDaemon(cfg aquareo.Config, db *bbolt.DB) *daemon {
	return &daemon{
		cfg: cfg,
		db:  db,
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
	ctrl := device.NewRPiController(s)

	if err := ctrl.Init(a.cfg); err != nil {
		return fmt.Errorf("daemon: Failed to start controller: %w", err)
	}

	if err := createBuckets(a.db, ctrl); err != nil {
		return fmt.Errorf("daemon: Failed to create buckets: %w", err)
	}

	a.ws = api.NewServer(a.cfg, ctrl)
	a.sc = sensor.NewCommander(a.cfg, ctrl)

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