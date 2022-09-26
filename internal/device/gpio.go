//go:build !arm
// +build !arm

package device

type gpio struct{}

func NewRPIODriver() *gpio {
	return &gpio{}
}

func (g *gpio) Open() error {
	return nil
}

func (g *gpio) Close() error {
	return nil
}
