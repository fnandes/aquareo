package store

import (
	"database/sql"
	"fmt"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"log"
	"os"
)

type store struct {
	db *sql.DB
}

func NewFileStore() *store {
	return &store{}
}

func (s *store) Store(bucket string, entry aquareo.MetricEntry) {
	f, err := os.OpenFile(fmt.Sprintf("dat_%s.txt", bucket), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("filestore: Failed to open data file", err.Error())
	}
	if _, err := f.Write([]byte(fmt.Sprintf("%v %v\n", entry.Timespan, entry.Value))); err != nil {
		log.Println("filestore: Failed to write data file", err.Error())
	}
	if err := f.Close(); err != nil {
		log.Println("filestore: Failed to close data file", err.Error())
	}
}

func (s *store) ReadAll(bucket string, size int) []aquareo.MetricEntry {
	panic("implement me")
}
