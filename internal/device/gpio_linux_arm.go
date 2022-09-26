//go:build linux && arm
// +build linux,arm

package device

import rpio "github.com/stianeikeland/go-rpio/v4"

type gpio struct{}

func NewRPIODriver() *gpio {
	return &gpio{}
}

func (g *gpio) Open() error {
	return rpio.Open()
}

func (g *gpio) Close() error {
	return rpio.Close()
}
