package sensor

import (
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/pedrobfernandes/aquareo/internal/device"
	"log"
	"time"
)

type cmd struct {
	config  aquareo.Config
	sensors []aquareo.Sensor
}

func NewCommander(config aquareo.Config) *cmd {
	return &cmd{
		config:  config,
		sensors: configureSensors(config),
	}
}

func (c *cmd) Start() {
	log.Println("starting sensors module ...")

	for {
		for _, s := range c.sensors {
			s.Collect()
		}
		time.Sleep(5 * time.Second)
	}
}

func configureSensors(config aquareo.Config) []aquareo.Sensor {
	var sensors []aquareo.Sensor

	for _, s := range config.Sensors {
		if s.Type == aquareo.DSL8B20 {
			sensors = append(sensors, device.NewDsl8b20Sensor(s.Id, s.Name))
		}
	}

	return sensors
}
