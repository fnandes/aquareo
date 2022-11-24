package daemon

import (
	"context"
	"encoding/json"

	"github.com/pedrobfernandes/aquareo/internal/api"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/pedrobfernandes/aquareo/internal/device"
	"github.com/pedrobfernandes/aquareo/internal/modules"
	"github.com/pedrobfernandes/aquareo/internal/store"
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

	a.ws = api.NewServer(a.ctrl)

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
