package store

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"math"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"go.etcd.io/bbolt"
)

const (
	TemperatureBucket = "temperature"
)

type storage struct {
	db *bbolt.DB
}

func NewBoldDbStorage(db *bbolt.DB) *storage {
	return &storage{db: db}
}

func (s *storage) CreateBucket(bucket string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("bbolt: failed to created bucket: %w", err)
		}
		return nil
	})
}

type objectStore struct {
	bucket string
	db     *bbolt.DB
}

func (s *storage) MetricStore(bucket string) aquareo.MetricStore {
	return &objectStore{
		db:     s.db,
		bucket: bucket,
	}
}

func (o *objectStore) List(size int) ([]aquareo.MetricEntry, error) {
	var arr []aquareo.MetricEntry

	err := o.db.View(func(tx *bbolt.Tx) error {
		cur := tx.Bucket([]byte(o.bucket)).Cursor()
		i := 0

		for k, v := cur.Last(); k != nil && i < size; k, v = cur.Prev() {
			arr = append(arr, aquareo.MetricEntry{
				Timespan: int64(binary.BigEndian.Uint64(k)),
				Value:    math.Float32frombits(binary.BigEndian.Uint32(v)),
			})
			i += 1
		}

		return nil
	})

	return arr, err
}

func (o *objectStore) Put(timespan int64, value float32) error {
	return o.db.Update(func(tx *bbolt.Tx) error {
		var kbuf, vbuf bytes.Buffer

		if err := binary.Write(&kbuf, binary.BigEndian, timespan); err != nil {
			return err
		}

		if err := binary.Write(&vbuf, binary.BigEndian, value); err != nil {
			return err
		}

		return tx.Bucket([]byte(o.bucket)).Put(kbuf.Bytes(), vbuf.Bytes())
	})
}
