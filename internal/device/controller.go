package device

import (
	"context"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/spf13/afero"
)

type controller struct {
	gpio       aquareo.GPIODriver
	fs         afero.Fs
	storage    aquareo.Storage
	cfgMgr     aquareo.Configurer
	subSystems []aquareo.Subsystem
}

func NewRPiController(
	fs afero.Fs,
	gpio aquareo.GPIODriver,
	storage aquareo.Storage,
	cfgMgr aquareo.Configurer,
) *controller {
	return &controller{
		gpio:    gpio,
		fs:      fs,
		storage: storage,
		cfgMgr:  cfgMgr,
	}
}

func (c *controller) Install(s aquareo.Subsystem) error {
	return s.Install(c)
}

func (c *controller) Start() error {
	if err := c.gpio.Open(); err != nil {
		return err
	}

	return nil
}

func (c *controller) Stop(ctx context.Context) {
	c.gpio.Close()
}

func (c *controller) Storage() aquareo.Storage {
	return c.storage
}

func (c *controller) Config() aquareo.Configurer {
	return c.cfgMgr
}
