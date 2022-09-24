package device

import (
	"errors"
	"strconv"
	"strings"

	"github.com/spf13/afero"
)

var (
	ErrReadSensor = errors.New("unable to read sensor data")
)

type ds18b20 struct {
	fs      afero.Fs
	id      string
	currVal float32
}

func NewDs18b20Sensor(fs afero.Fs, id string) *ds18b20 {
	return &ds18b20{
		fs: fs,
		id: id,
	}
}

func (s *ds18b20) Id() string {
	return s.id
}

func (s *ds18b20) CurrentValue() float32 {
	return s.currVal
}

func (s *ds18b20) Refresh() error {
	data, err := afero.ReadFile(s.fs, "/sys/bus/w1/devices/"+s.id+"/w1_slave")
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
