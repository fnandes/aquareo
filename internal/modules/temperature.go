package modules

import (
	"context"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/pedrobfernandes/aquareo/internal/store"
	"github.com/spf13/afero"
)

type module struct {
	cfg     aquareo.TemperatureControllerConfig
	store   aquareo.MetricStore
	fs      afero.Fs
	stopper chan struct{}
}

func NewTemperatureController(fs afero.Fs, cfg aquareo.TemperatureControllerConfig) *module {
	return &module{
		cfg:     cfg,
		fs:      fs,
		stopper: make(chan struct{}),
	}
}

func (tc *module) Install(ctrl aquareo.Controller) error {
	if err := ctrl.Storage().CreateBucket(store.TemperatureBucket); err != nil {
		return err
	}

	tc.store = ctrl.Storage().MetricStore(store.TemperatureBucket)

	return nil
}

func (tc *module) Start() {
	log.Println("temperature: module started")

	var wg sync.WaitGroup
	wg.Add(1)

	tick := time.NewTicker(time.Duration(tc.cfg.TickInterval) * time.Millisecond)

	go func() {
		defer wg.Done()

		for {
			select {
			case <-tc.stopper:
				log.Println("temperature: stopping")
				return
			case <-tick.C:
				val, err := tc.getValue()
				if err != nil {
					log.Printf("temperature: failed to get sensor data: %v\n", err)
				} else {
					tc.store.Put(time.Now().UTC().Unix(), val)
				}
			default:
				continue
			}
		}
	}()

	wg.Wait()
}

func (tc *module) Stop(ctx context.Context) {
	close(tc.stopper)
}

func (tc *module) getValue() (float32, error) {
	data, err := afero.ReadFile(tc.fs, "/sys/bus/w1/devices/"+tc.cfg.DeviceId+"/w1_slave")
	if err != nil {
		return -1, err
	}

	raw := string(data)

	if !strings.Contains(raw, " YES") {
		return -1, err
	}

	i := strings.LastIndex(raw, "t=")
	if i == -1 {
		return -1, err
	}

	c, err := strconv.ParseFloat(raw[i+2:len(raw)-1], 64)
	if err != nil {
		return -1, err
	}

	return float32(c / 1000.0), nil
}
