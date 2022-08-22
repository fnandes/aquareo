package store

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"math"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"go.etcd.io/bbolt"
)

type store struct {
	db *bbolt.DB
}

func NewBoltDbStore(db *bbolt.DB) *store {
	return &store{db: db}
}

func (s *store) CreateBucketIfNotExists(bucket string) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			return fmt.Errorf("bbolt: Failed to created bucket: %w", err)
		}
		return nil
	})
}

func (s *store) Store(bucket string, entry aquareo.MetricEntry) error {
	log.Printf("bbolt: Saving %s: [%v, %v]", bucket, entry.Timespan, entry.Value)
	return s.db.Update(func(tx *bbolt.Tx) error {
		var kbuf, vbuf bytes.Buffer

		if err := binary.Write(&kbuf, binary.LittleEndian, entry.Timespan); err != nil {
			return fmt.Errorf("bbolt: Failed to get bytes from key: %w", err)
		}

		if err := binary.Write(&vbuf, binary.LittleEndian, entry.Value); err != nil {
			return fmt.Errorf("bbolt: Failed to get bytes from value: %w", err)
		}

		return tx.Bucket([]byte(bucket)).Put(kbuf.Bytes(), vbuf.Bytes())
	})
}

func (s *store) ReadAll(bucket string, size int) ([]aquareo.MetricEntry, error) {
	var arr []aquareo.MetricEntry

	err := s.db.View(func(tx *bbolt.Tx) error {
		cur := tx.Bucket([]byte(bucket)).Cursor()
		i := 0

		for k, v := cur.Last(); k != nil && i < size; k, v = cur.Prev() {
			arr = append(arr, aquareo.MetricEntry{
				Timespan: int64(binary.LittleEndian.Uint64(k)),
				Value:    math.Float32frombits(binary.LittleEndian.Uint32(v)),
			})
			i += 1
		}

		return nil
	})

	return arr, err
}
