package device

import (
	"errors"
	"io/ioutil"
	"strconv"
	"strings"
)

var (
	ErrReadSensor = errors.New("unable to read sensor data")
)

type ds18b20 struct {
	id       string
	serialNr string
	currVal  float32
}

func NewDs18b20Sensor(id string, serialNr string) *ds18b20 {
	return &ds18b20{
		id:       id,
		serialNr: serialNr,
	}
}

func (s *ds18b20) Id() string {
	return s.id
}

func (s *ds18b20) CurrentValue() float32 {
	return s.currVal
}

func (s *ds18b20) Refresh() error {
	data, err := ioutil.ReadFile("/sys/bus/w1/devices/" + s.id + "/w1_slave")
	if err != nil {
		return ErrReadSensor
	}

	raw := string(data)

	if !strings.Contains(raw, " YES") {
		return ErrReadSensor
	}

	i := strings.LastIndex(raw, "t=")
	if i == -1 {
		return ErrReadSensor
	}

	c, err := strconv.ParseFloat(raw[i+2:len(raw)-1], 64)
	if err != nil {
		return ErrReadSensor
	}

	s.currVal = float32(c / 1000.0)

	return nil
}
