package device

import (
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

func (c *controller) Open() error {
	return rpio.Open()
	//return nil
}

func (c *controller) Close() error {
	return rpio.Close()
}

func (c *controller) Store() aquareo.Store {
	return c.store
}

func (c *controller) SetInput(pin uint8) {
	rpio.Pin(pin).Input()
}

func (c *controller) SetOutput(pin uint8) {
	rpio.Pin(pin).Output()
}

func (c *controller) High(pin uint8) {
	rpio.Pin(pin).High()
}

func (c *controller) Low(pin uint8) {
	rpio.Pin(pin).Low()
}

func (c *controller) AddSensor(id string, sensor aquareo.Sensor) error {
	c.sensors[id] = sensor
	return nil
}

func (c *controller) Sensor(id string) aquareo.Sensor {
	return c.sensors[id]
}
