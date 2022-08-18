package device

import (
	"fmt"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/stianeikeland/go-rpio"
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

func (c *controller) Init(conf aquareo.Config) error {
	if err := rpio.Open(); err != nil {
		return fmt.Errorf("controller: Failed to open the GPIO access: %w", err)
	}

	// register all sensors
	for _, s := range conf.Sensors {
		if s.Type == aquareo.DSL8B20 {
			c.sensors[s.Id] = NewDsl8b20Sensor(s.Id, s.Name)
		}
	}

	// add the system temperature sensor
	c.sensors[aquareo.SensorSysTemp] = NewSysTempSensor(aquareo.SensorSysTemp, "Controller Temperature")

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

func (c *controller) RegisterSensor(sensor aquareo.Sensor) error {
	if err := c.store.CreateBucketIfNotExists(sensor.Id()); err != nil {
		return fmt.Errorf("controller: Failed to create store bucket for sensor %s: %w", sensor.Id(), err)
	}

	c.sensors[sensor.Id()] = sensor
	return nil
}

func (c *controller) Sensor(id string) aquareo.Sensor {
	return c.sensors[id]
}
