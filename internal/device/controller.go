package device

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/stianeikeland/go-rpio"
)

var (
	ErrTempSensorNotFound = errors.New("unable to find a temperature sensor")
)

type controller struct {
	store   aquareo.Store
	sensors map[string]aquareo.Sensor
}

func NewRPiController(store aquareo.Store) *controller {
	return &controller{
		store:   store,
		sensors: make(map[string]aquareo.Sensor),
	}
}

func (c *controller) Init(cfg aquareo.Config) error {
	if err := rpio.Open(); err != nil {
		return err
	}

	// add the system temperature sensor
	c.sensors[aquareo.SensorSysTemp] = NewSysTempSensor(aquareo.SensorSysTemp, "Controller Temperature")

	// detect and install the temperature sensors
	data, err := ioutil.ReadFile("/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves")
	if err != nil {
		return ErrTempSensorNotFound
	}

	for _, sid := range strings.Split(string(data), "\n") {
		if sid != "" {
			log.Println("controller: installing sensor", sid)
			c.sensors[sid] = NewDs18b20Sensor(sid, sid)
		}
	}

	return nil
}

func (c *controller) Close() error {
	return rpio.Close()
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
