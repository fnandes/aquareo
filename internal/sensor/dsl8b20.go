package sensor

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
	fs afero.Fs
	id string
}

func NewDs18b20Sensor(fs afero.Fs, id string) *ds18b20 {
	return &ds18b20{
		fs: fs,
		id: id,
	}
}

func (s *ds18b20) GetValue() (float32, error) {
	data, err := afero.ReadFile(s.fs, "/sys/bus/w1/devices/"+s.id+"/w1_slave")
	if err != nil {
		return -1, ErrReadSensor
	}

	raw := string(data)

	if !strings.Contains(raw, " YES") {
		return -1, ErrReadSensor
	}

	i := strings.LastIndex(raw, "t=")
	if i == -1 {
		return -1, ErrReadSensor
	}

	c, err := strconv.ParseFloat(raw[i+2:len(raw)-1], 64)
	if err != nil {
		return -1, ErrReadSensor
	}

	return float32(c / 1000.0), nil
}
