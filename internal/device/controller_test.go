package device_test

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/pedrobfernandes/aquareo/internal/device"
	"github.com/pedrobfernandes/aquareo/mocks"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestRpiController(t *testing.T) {
	ctrl := gomock.NewController(t)
	fs := afero.NewMemMapFs()
	gpio := mocks.NewMockGPIODriver(ctrl)
	storage := mocks.NewMockStorage(ctrl)
	installable := mocks.NewMockSubsystem(ctrl)

	deviceCtrl := device.NewRPiController(fs, gpio, storage)

	t.Run("Install", func(t *testing.T) {
		installable.EXPECT().
			Install(deviceCtrl).
			Times(1)
		assert.NoError(t, deviceCtrl.Install(installable))
	})

	t.Run("Start", func(t *testing.T) {
		installable.EXPECT().Start()
		gpio.EXPECT().Open().Return(nil)

		assert.NoError(t, deviceCtrl.Start())
	})

	t.Run("Stop", func(t *testing.T) {
		gpio.EXPECT().Close().Return(nil)
		installable.EXPECT().Stop(context.TODO())

		deviceCtrl.Stop(context.TODO())
	})
}
