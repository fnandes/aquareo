package device

import (
	"os"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/pedrobfernandes/aquareo/mocks"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestInit_WithTempSensors(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	fs := afero.NewMemMapFs()
	store := mocks.NewMockStore(ctrl)
	driver := mocks.NewMockGPIODriver(ctrl)

	cfg := new(aquareo.Config)
	c := NewRPiController(fs, driver, store)

	driver.EXPECT().Open().Return(nil)
	afero.WriteFile(fs, "/sys/bus/w1/devices/w1_bus_master1/w1_master_slaves", []byte("28-aaabbbb\n28-bbbcccc"), os.ModeAppend)

	c.Init(*cfg)

	assert.Len(t, c.sensors, 2)
	assert.IsType(t, c.sensors["28-aaabbbb"], NewDs18b20Sensor(fs, "28-aaabbbb"))
	assert.IsType(t, c.sensors["28-bbbcccc"], NewDs18b20Sensor(fs, "28-bbbcccc"))
}
