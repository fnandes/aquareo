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
	subSystems []aquareo.Subsystem
}

func NewRPiController(
	fs afero.Fs,
	gpio aquareo.GPIODriver,
	storage aquareo.Storage,
) *controller {
	return &controller{
		gpio:    gpio,
		fs:      fs,
		storage: storage,
	}
}

func (c *controller) Install(s aquareo.Subsystem) error {
	c.subSystems = append(c.subSystems, s)
	return s.Install(c)
}

func (c *controller) Start() error {
	if err := c.gpio.Open(); err != nil {
		return err
	}

	for _, ss := range c.subSystems {
		go ss.Start()
	}

	return nil
}

func (c *controller) Stop(ctx context.Context) {
	c.gpio.Close()

	for _, ss := range c.subSystems {
		ss.Stop(ctx)
	}
}

func (c *controller) Storage() aquareo.Storage {
	return c.storage
}
