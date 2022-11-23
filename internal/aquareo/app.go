package aquareo

import (
	"context"
)

//go:generate mockgen -source=app.go -destination=../../mocks/app_mocks.go -package=mocks Controller,Storage,MetricStore,Subsystem,GPIODriver

type WebServer interface {
	Start(addr string)
	Stop(ctx context.Context)
}

type Controller interface {
	Install(s Subsystem) error
	Start() error
	Stop(ctx context.Context)

	Storage() Storage
}

type GPIODriver interface {
	Open() error
	Close() error
}

type Subsystem interface {
	Install(ctrl Controller) error
	Start()
	Stop(ctx context.Context)
}

type Storage interface {
	MetricStore(bucket string) MetricStore
	CreateBucket(bucket string) error
}

type MetricStore interface {
	Put(timespan int64, value float32) error
	List(size int) ([]MetricEntry, error)
}

type MetricEntry struct {
	Timespan int64
	Value    float32
}
