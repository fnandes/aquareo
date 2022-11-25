package store_test

import (
	"os"
	"testing"
	"time"

	"github.com/fnandes/aquareo/internal/aquareo"
	"github.com/fnandes/aquareo/internal/store"
	"github.com/stretchr/testify/assert"
	"go.etcd.io/bbolt"
)

func TestStorage_CreateBucket(t *testing.T) {
	storage, dbPath := MustCreateStorage()
	defer os.Remove(dbPath)

	storage.CreateBucket("newbucket")

	// create bucket
	assert.NoError(t, storage.CreateBucket("newbucket"))

	// create bucket again
	assert.NoError(t, storage.CreateBucket("newbucket"))
}

func TestStorage_MetricStore(t *testing.T) {
	storage, dbPath := MustCreateStorage()
	defer os.Remove(dbPath)

	metricStore := storage.MetricStore("temp")
	assert.NoError(t, storage.CreateBucket("temp"))

	oldest := aquareo.MetricEntry{Timespan: time.Now().Add(-2 * time.Second).Unix(), Value: float32(23.0002)}
	previous := aquareo.MetricEntry{Timespan: time.Now().Add(-1 * time.Second).Unix(), Value: float32(2300)}
	recent := aquareo.MetricEntry{Timespan: time.Now().Unix(), Value: float32(0.000021)}

	for _, entry := range []aquareo.MetricEntry{oldest, previous, recent} {
		assert.NoError(t, metricStore.Put(entry.Timespan, entry.Value))
	}

	list, err := metricStore.List(2)

	assert.NoError(t, err)
	assert.Equal(t, 2, len(list))

	assert.Equal(t, recent, list[0])
	assert.Equal(t, previous, list[1])
}

func MustCreateStorage() (aquareo.Storage, string) {
	path := createTempDbFile()
	db, err := bbolt.Open(path, 0600, nil)
	if err != nil {
		panic(err)
	}

	storage := store.NewBoldDbStorage(db)

	return storage, path
}

func createTempDbFile() string {
	file, err := os.CreateTemp("", "aq-test-")
	if err != nil {
		panic(err)
	}

	return file.Name()
}
