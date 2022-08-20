package sensor

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
)

const (
	RefreshThresholdSecs = 2
	StoreThresholdSecs   = 60
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
	var wg sync.WaitGroup

	go c.schedule(&wg, RefreshThresholdSecs*time.Second, func() {
		for _, s := range c.ctrl.Sensors() {
			if err := s.Refresh(); err != nil {
				log.Printf("sensors: Failed to refresh: %s: %v\n", s.Id(), err.Error())
			}
		}
	})

	go c.schedule(&wg, StoreThresholdSecs*time.Second, func() {
		for _, s := range c.ctrl.Sensors() {
			err := c.ctrl.Store().Store(s.Id(), aquareo.MetricEntry{
				Timespan: time.Now().UTC().Unix(),
				Value:    s.CurrentValue(),
			})
			if err != nil {
				log.Printf("sensors: Failed to store %s: %v\n", s.Id(), err.Error())
			}
		}
	})

	wg.Wait()
}

func (c *cmd) schedule(wg *sync.WaitGroup, threshold time.Duration, task func()) {
	wg.Add(1)
	defer wg.Done()

	for {
		time.Sleep(threshold)

		select {
		case <-c.stopped:
			return
		default:
		}
		task()
	}
}

func (c *cmd) Stop(ctx context.Context) {
	close(c.stopped)
	log.Println("sensors: Module closed")
}
