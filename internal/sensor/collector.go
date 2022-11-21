package sensor

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/spf13/afero"
)

const (
	RefreshThresholdSecs = 2
	StoreThresholdSecs   = 60
)

type cmd struct {
	ctrl    aquareo.Controller
	fs      afero.Fs
	stopped chan struct{}
}

func NewDataCollector(c aquareo.Controller, fs afero.Fs) *cmd {
	return &cmd{
		ctrl:    c,
		stopped: make(chan struct{}),
		fs:      fs,
	}
}

func (c *cmd) Start() {
	log.Println("sensors: Module started")

	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-c.stopped:
				return
			default:
			}

			cfg, err := c.ctrl.Config().Get()
			if err != nil {
				log.Printf("sensors: unable to read config: %v", err)
				return
			}

			for _, sc := range cfg.Sensors {
				if err := c.collectSensorData(sc); err != nil {
					log.Printf("sensors: unable to read sensor data for %s: %v", sc.Id, err)
					continue
				}
			}

			time.Sleep(RefreshThresholdSecs * time.Second)
		}
	}()

	wg.Wait()
}

func (c *cmd) Stop(ctx context.Context) {
	close(c.stopped)
	log.Println("sensors: Module closed")
}

func (c *cmd) collectSensorData(sc aquareo.SensorConfig) error {
	var sensor aquareo.Sensor
	if sc.Type == aquareo.DSL8B20 {
		sensor = NewDs18b20Sensor(c.fs, sc.Id)
	}

	val, err := sensor.GetValue()
	if err != nil {
		return fmt.Errorf("collectSensorData: failed to get value for sensor %s: %v", sc.Id, err)
	}

	return c.ctrl.Store().Store(sc.Id, aquareo.MetricEntry{
		Timespan: time.Now().UTC().Unix(),
		Value:    val,
	})
}
