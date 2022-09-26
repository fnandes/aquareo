package device

import (
	"errors"
	"log"
	"strings"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/spf13/afero"
)

var (
	ErrTempSensorNotFound = errors.New("unable to find a temperature sensor")
)

type controller struct {
	gpio    aquareo.GPIODriver
	fs      afero.Fs
	store   aquareo.Store
	sensors map[string]aquareo.Sensor
}

func NewRPiController(fs afero.Fs, gpio aquareo.GPIODriver, store aquareo.Store) *controller {
	return &controller{
		gpio:    gpio,
		fs:      fs,
		store:   store,
		sensors: make(map[string]aquareo.Sensor),
	}
}

func (c *controller) Init(cfg aquareo.Config) error {
	if err := c.gpio.Open(); err != nil {
		return err
	}

	// detect and install the temperature sensors
	data, err := afero.ReadFile(c.fs, "/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves")
	if err != nil {
		return ErrTempSensorNotFound
	}

	for _, sid := range strings.Split(string(data), "\n") {
		if sid != "" {
			log.Println("controller: installing sensor", sid)
			c.sensors[sid] = NewDs18b20Sensor(c.fs, sid)
		}
	}

	return nil
}

func (c *controller) Close() error {
	return c.gpio.Close()
}

func (c *controller) Store() aquareo.Store {
	return c.store
}

func (c *controller) Sensors() []aquareo.Sensor {
	var arr []aquareo.Sensor

	for _, s := range c.sensors {
		arr = append(arr, s)
	}
	return arr
}

func (c *controller) Sensor(id string) aquareo.Sensor {
	return c.sensors[id]
}
