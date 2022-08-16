package device

type dsl8b20 struct {
	id       string
	serialNr string
}

func NewDsl8b20Sensor(id string, serialNr string) *dsl8b20 {
	return &dsl8b20{
		id:       id,
		serialNr: serialNr,
	}
}

func (s *dsl8b20) Id() string {
	return s.id
}

func (s *dsl8b20) CurrentValue() float32 {
	return 2.5
}

func (s *dsl8b20) Refresh() error {
	return nil
}
