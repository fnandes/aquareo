package sensor

import (
	"context"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/pedrobfernandes/aquareo/internal/device"
	"log"
	"time"
)

type cmd struct {
	ctrl    aquareo.Controller
	config  aquareo.Config
	sensors []aquareo.Sensor
	stopped chan struct{}
}

func NewCommander(config aquareo.Config, c aquareo.Controller) *cmd {
	return &cmd{
		config:  config,
		ctrl:    c,
		sensors: configureSensors(config),
		stopped: make(chan struct{}),
	}
}

func (c *cmd) Start() {
	log.Println("sensors: Module started")

	for {
		select {
		case <-c.stopped:
			return
		default:
		}

		for _, s := range c.sensors {
			if err := s.Refresh(); err != nil {
				log.Printf("sensors: Failed to refresh: %s: %v\n", s.Id(), err.Error())
			}

			c.ctrl.Store().Store(s.Id(), aquareo.MetricEntry{
				Timespan: int(time.Now().UTC().Unix()),
				Value:    s.CurrentValue(),
			})
		}

		time.Sleep(5 * time.Second)
	}
}

func (c *cmd) Stop(ctx context.Context) {
	close(c.stopped)
	log.Println("sensors: Module closed")
}

func configureSensors(config aquareo.Config) []aquareo.Sensor {
	var sensors []aquareo.Sensor

	for _, s := range config.Sensors {
		if s.Type == aquareo.DSL8B20 {
			sensors = append(sensors, device.NewDsl8b20Sensor(s.Id, s.Name))
		}
	}

	// add the system temperature sensor
	sensors = append(sensors, device.NewSysTempSensor(aquareo.SensorSysTemp, "Controller Temperature"))

	return sensors
}
