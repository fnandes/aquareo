package store

import (
	"bytes"
	"encoding/binary"
	"fmt"

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

func (s *store) ReadAll(bucket string, size int) []aquareo.MetricEntry {
	panic("implement me")
}
