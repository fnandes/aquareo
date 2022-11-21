package device

import (
	"errors"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/spf13/afero"
)

var (
	ErrTempSensorNotFound = errors.New("unable to find a temperature sensor")
)

type controller struct {
	gpio       aquareo.GPIODriver
	fs         afero.Fs
	store      aquareo.Store
	configurer aquareo.Configurer
}

func NewRPiController(
	fs afero.Fs,
	gpio aquareo.GPIODriver,
	configurer aquareo.Configurer,
	store aquareo.Store,
) *controller {
	return &controller{
		gpio:       gpio,
		fs:         fs,
		store:      store,
		configurer: configurer,
	}
}

func (c *controller) Config() aquareo.Configurer {
	return c.configurer
}

func (c *controller) Init() error {
	if err := c.gpio.Open(); err != nil {
		return err
	}

	return nil
}

func (c *controller) Close() error {
	return c.gpio.Close()
}

func (c *controller) Store() aquareo.Store {
	return c.store
}
