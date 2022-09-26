package aquareo

import (
	"context"
)

//go:generate mockgen -source=app.go -destination=../../mocks/app_mocks.go -package=mocks Controller,Sensor,SensorCommander,Store,GPIODriver

type WebServer interface {
	Start(addr string)
	Stop(ctx context.Context)
}

type Controller interface {
	Store() Store
	Sensors() []Sensor
	Sensor(sensorId string) Sensor
}

type GPIODriver interface {
	Open() error
	Close() error
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
	Store(bucket string, entry MetricEntry) error
	ReadAll(bucket string, size int) ([]MetricEntry, error)
}

type MetricEntry struct {
	Timespan int64
	Value    float32
}
