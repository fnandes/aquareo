package aquareo

import (
	"context"
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

func (a *App) Stop(ctx context.Context) {
	a.ws.Stop(ctx)
	a.sc.Stop(ctx)
}

func (a *App) Start() {
	go func() {
		a.ws.Start(":8082")
	}()

	go func() {
		a.sc.Start()
	}()
}

type WebServer interface {
	Start(addr string)
	Stop(ctx context.Context)
}

type Controller interface {
	Store() Store

	SetInput(pin uint8)
	SetOutput(pin uint8)
	High(pin uint8)
	Low(pin uint8)

	RegisterSensor(sensor Sensor) error
	Sensors() []Sensor
	Sensor(sensorId string) Sensor
}

type Sensor interface {
	Id() string
	CurrentValue() float32
	Refresh() error
}

type SensorCommander interface {
	Start()
	Stop(ctx context.Context)
}

type Store interface {
	CreateBucketIfNotExists(bucket string) error
	Store(bucket string, entry MetricEntry) error
	ReadAll(bucket string, size int) []MetricEntry
}
type MetricEntry struct {
	Timespan int
	Value    float32
}
