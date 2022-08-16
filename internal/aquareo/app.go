package aquareo

import (
	"sync"
)

type App struct {
	config Config
	c      Controller
	ws     WebServer
	sc     SensorCommander
}

func NewApp(config Config, c Controller, server WebServer, commander SensorCommander) *App {
	return &App{
		config: config,
		c:      c,
		ws:     server,
		sc:     commander,
	}
}

func (a *App) Start(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		a.ws.Start(":8080")
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		a.sc.Start()
	}()
}

type WebServer interface {
	Start(addr string)
}

type Controller interface {
	Store() Store

	SetInput(pin uint8)
	SetOutput(pin uint8)
	High(pin uint8)
	Low(pin uint8)

	AddSensor(sensorId string, sensor Sensor) error
	Sensor(sensorId string) Sensor
}

type Sensor interface {
	CurrentValue()
	Collect()
}

type SensorCommander interface {
	Start()
}

type Store interface {
	Store(bucket string, entry MetricEntry)
	ReadAll(bucket string, size int) []MetricEntry
}
type MetricEntry struct {
	Timespan int
	Value    float32
}
