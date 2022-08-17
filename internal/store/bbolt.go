package store

import (
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"go.etcd.io/bbolt"
	"math"
)

type store struct {
	db *bbolt.DB
}

func NewBoltDbStore(db *bbolt.DB) *store {
	return &store{db: db}
}

func (s *store) Store(bucket string, entry aquareo.MetricEntry) error {
	return s.db.Update(func(tx *bbolt.Tx) error {
		return tx.Bucket([]byte(bucket)).Put([]byte(entry.Timespan), []byte(entry.Value))
	})
}

func (s *store) ReadAll(bucket string, size int) []aquareo.MetricEntry {
	panic("implement me")
}
