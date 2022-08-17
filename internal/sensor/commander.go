package sensor

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/pedrobfernandes/aquareo/internal/device"
)

type cmd struct {
	ctrl    aquareo.Controller
	config  aquareo.Config
	stopped chan struct{}
}

func NewCommander(config aquareo.Config, c aquareo.Controller) *cmd {
	return &cmd{
		config:  config,
		ctrl:    c,
		stopped: make(chan struct{}),
	}
}

func (c *cmd) Start() {
	log.Println("sensors: Module started")

	if err := c.configureSensors(); err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case <-c.stopped:
			return
		default:
		}

		for _, s := range c.ctrl.Sensors() {
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

func (c *cmd) configureSensors() error {
	for _, s := range c.config.Sensors {
		if s.Type == aquareo.DSL8B20 {
			if err := c.ctrl.RegisterSensor(device.NewDsl8b20Sensor(s.Id, s.Name)); err != nil {
				return fmt.Errorf("sensor commander: Failed to register sensor: %w", err)
			}
		}
	}

	// add the system temperature sensor
	if err := c.ctrl.RegisterSensor(device.NewSysTempSensor(aquareo.SensorSysTemp, "Controller Temperature")); err != nil {
		return fmt.Errorf("sensor commander: Failed to register sensor: %w", err)
	}

	return nil
}
