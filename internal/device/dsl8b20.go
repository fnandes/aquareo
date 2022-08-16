package device

import "log"

type dsl8b20 struct {
	id       string
	serialNr string
}

func (s dsl8b20) CurrentValue() {
	panic("implement me")
}

func (s dsl8b20) Collect() {
	log.Printf("collecting %s\n", s.id)
}

func NewDsl8b20Sensor(id string, serialNr string) *dsl8b20 {
	return &dsl8b20{
		id:       id,
		serialNr: serialNr,
	}
}
