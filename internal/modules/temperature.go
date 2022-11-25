package modules

import (
	"context"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/fnandes/aquareo/internal/aquareo"
	"github.com/fnandes/aquareo/internal/store"
	"github.com/spf13/afero"
)

type module struct {
	deviceId         string
	tickInterval     int32
	snapshotInterval int32
	store            aquareo.MetricStore
	fs               afero.Fs
	stopper          chan struct{}
}

func NewTemperatureController(fs afero.Fs, cfg aquareo.TemperatureControllerConfig) *module {
	return &module{
		deviceId:         cfg.DeviceId,
		tickInterval:     cfg.TickInterval,
		snapshotInterval: cfg.SnapshotInterval,
		fs:               fs,
		stopper:          make(chan struct{}),
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

	tick := time.NewTicker(time.Duration(tc.tickInterval) * time.Millisecond)
	snapshot := time.NewTicker(time.Duration(tc.snapshotInterval) * time.Millisecond)

	go func() {
		defer wg.Done()

		var currVal float32
		for {
			select {
			case <-tick.C:
				val, err := tc.getValue()
				if err != nil {
					log.Println(err)
				} else {
					currVal = val
				}
			case <-snapshot.C:
				if err := tc.store.Put(time.Now().UTC().Unix(), currVal); err != nil {
					log.Println(err)
				}
			case <-tc.stopper:
				log.Println("temperature: stopping")
				return
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
	data, err := afero.ReadFile(tc.fs, "/sys/bus/w1/devices/"+tc.deviceId+"/w1_slave")
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
