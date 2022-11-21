package aquareo

import (
	"context"
)

//go:generate mockgen -source=app.go -destination=../../mocks/app_mocks.go -package=mocks Controller,Configurer,Sensor,DataCollector,Store,GPIODriver

type Configurer interface {
	Get() (Config, error)
	Save(cfg Config) error
}

type WebServer interface {
	Start(addr string)
	Stop(ctx context.Context)
}

type Controller interface {
	Store() Store
	Config() Configurer
}

type GPIODriver interface {
	Open() error
	Close() error
}

type Sensor interface {
	GetValue() (float32, error)
}

type DataCollector interface {
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
