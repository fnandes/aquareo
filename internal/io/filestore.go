package io

import "github.com/pedrobfernandes/aquareo/internal/aquareo"

type store struct {
}

func (s *store) Store(bucket string, entry aquareo.MetricEntry) {
	panic("implement me")
}

func (s *store) ReadAll(bucket string, size int) []aquareo.MetricEntry {
	panic("implement me")
}

func NewFileStorage() *store {
	return &store{}
}
