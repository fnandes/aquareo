package sensor

import (
	"context"
	"log"
	"time"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
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

			err := c.ctrl.Store().Store(s.Id(), aquareo.MetricEntry{
				Timespan: int(time.Now().UTC().Unix()),
				Value:    s.CurrentValue(),
			})
			if err != nil {
				log.Println("sensors: Failed to store: ", err.Error())
			}
		}

		time.Sleep(5 * time.Second)
	}
}

func (c *cmd) Stop(ctx context.Context) {
	close(c.stopped)
	log.Println("sensors: Module closed")
}
