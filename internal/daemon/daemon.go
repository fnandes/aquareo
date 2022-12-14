package daemon

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/fnandes/aquareo/internal/api"
	"github.com/fnandes/aquareo/internal/aquareo"
	"github.com/fnandes/aquareo/internal/device"
	"github.com/fnandes/aquareo/internal/modules"
	"github.com/fnandes/aquareo/internal/store"
	"github.com/spf13/afero"
	"go.etcd.io/bbolt"
)

type daemon struct {
	fs   afero.Fs
	db   *bbolt.DB
	ws   aquareo.WebServer
	ctrl aquareo.Controller
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

	if a.ctrl != nil {
		a.ctrl.Stop(ctx)
	}
}

func (a *daemon) Start() error {
	s := store.NewBoldDbStorage(a.db)
	gpio := device.NewRPIODriver()

	cfg, err := loadConfig(a.fs, "config.json")
	if err != nil {
		panic(err)
	}

	a.ctrl = device.NewRPiController(a.fs, gpio, s, cfg)

	if cfg.TemperatureController.Enabled {
		tc := modules.NewTemperatureController(a.fs, cfg.TemperatureController)
		a.ctrl.Install(tc)
	}

	for _, cm := range cfg.CustomMetrics {
		if err := s.CreateBucket(fmt.Sprintf("cm_%s", cm)); err != nil {
			log.Fatal(err)
		}
	}

	a.ws = api.NewServer(a.ctrl, cfg)

	go func() {
		a.ctrl.Start()
	}()

	go func() {
		a.ws.Start(":8082")
	}()

	return nil
}

func loadConfig(fs afero.Fs, name string) (aquareo.Config, error) {
	buf, err := afero.ReadFile(fs, name)
	if err != nil {
		return aquareo.Config{}, err
	}

	var config aquareo.Config
	if err := json.Unmarshal(buf, &config); err != nil {
		return aquareo.Config{}, err
	}

	return config, nil
}
