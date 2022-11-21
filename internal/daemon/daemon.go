package daemon

import (
	"context"
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
	dc aquareo.DataCollector
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
	if a.dc != nil {
		a.dc.Stop(ctx)
	}
}

func (a *daemon) Start() error {
	s := store.NewBoltDbStore(a.db)
	gpio := device.NewRPIODriver()
	cfg := store.NewFileConfigurer("config.json", a.fs)

	ctrl := device.NewRPiController(a.fs, gpio, cfg, s)

	if err := createBuckets(a.db, ctrl); err != nil {
		return fmt.Errorf("daemon: Failed to create buckets: %w", err)
	}

	a.ws = api.NewServer(ctrl)
	a.dc = sensor.NewDataCollector(ctrl, a.fs)

	go func() {
		a.ws.Start(":8082")
	}()

	go func() {
		a.dc.Start()
	}()
	return nil
}

func createBuckets(db *bbolt.DB, c aquareo.Controller) error {
	cfg, err := c.Config().Get()
	if err != nil {
		return err
	}
	return db.Update(func(tx *bbolt.Tx) error {
		for sid := range cfg.Sensors {
			if _, err := tx.CreateBucketIfNotExists([]byte(sid)); err != nil {
				return err
			}
		}
		return nil
	})
}
