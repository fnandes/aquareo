package device

import (
	"os/exec"
	"strconv"
	"strings"
)

type sTemp struct {
	id      string
	name    string
	currVal float32
}

func NewSysTempSensor(id string, name string) *sTemp {
	return &sTemp{
		id:   id,
		name: name,
	}
}

func (s *sTemp) Id() string {
	return s.id
}

func (s *sTemp) CurrentValue() float32 {
	return s.currVal
}

func (s *sTemp) Refresh() error {
	raw, err := exec.Command("vcgencmd", "measure_temp").Output()
	if err != nil {
		return err
	}

	str := strings.ReplaceAll(strings.ReplaceAll(string(raw), "temp=", ""), "'C\n", "")

	val, err := strconv.ParseFloat(str, 32)
	if err != nil {
		return err
	}
	s.currVal = float32(val)

	return nil
}
