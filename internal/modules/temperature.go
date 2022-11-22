package modules

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/pedrobfernandes/aquareo/internal/store"
	"github.com/spf13/afero"
)

const (
	RefreshThresholdSecs = 2
	StoreThresholdSecs   = 60
)

var (
	ErrReadSensor = errors.New("unable to read sensor data")
)

type module struct {
	deviceId string
	store    aquareo.MetricStore
	fs       afero.Fs
	stopper  chan struct{}
}

func NewTemperatureController(deviceId string, fs afero.Fs) *module {
	return &module{
		deviceId: deviceId,
		fs:       fs,
		stopper:  make(chan struct{}),
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

	go func() {
		defer wg.Done()

		for {
			select {
			case <-tc.stopper:
				return
			default:
			}

			val, err := tc.getValue()
			if err != nil {
				log.Printf("temperature: failed to get sensor data: %v\n", err)
			} else {
				tc.store.Put(time.Now().UTC().Unix(), val)
			}

			time.Sleep(RefreshThresholdSecs * time.Second)
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
		return -1, ErrReadSensor
	}

	raw := string(data)

	if !strings.Contains(raw, " YES") {
		return -1, ErrReadSensor
	}

	i := strings.LastIndex(raw, "t=")
	if i == -1 {
		return -1, ErrReadSensor
	}

	c, err := strconv.ParseFloat(raw[i+2:len(raw)-1], 64)
	if err != nil {
		return -1, ErrReadSensor
	}

	return float32(c / 1000.0), nil
}
