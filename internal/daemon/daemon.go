package daemon

import (
	"context"

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

	tc := modules.NewTemperatureController("device-id", a.fs)
	cfgMgr := store.NewFileConfigurer("config.json", a.fs)

	a.ctrl = device.NewRPiController(a.fs, gpio, s, cfgMgr)
	a.ctrl.Install(tc)

	a.ws = api.NewServer(a.ctrl)

	go func() {
		a.ctrl.Start()
	}()

	go func() {
		a.ws.Start(":8082")
	}()

	return nil
}
