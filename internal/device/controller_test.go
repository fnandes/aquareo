package device_test

import (
	"context"
	"testing"

	"github.com/fnandes/aquareo/internal/aquareo"
	"github.com/fnandes/aquareo/internal/device"
	"github.com/fnandes/aquareo/mocks"
	"github.com/golang/mock/gomock"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"
)

func TestRpiController(t *testing.T) {
	ctrl := gomock.NewController(t)
	fs := afero.NewMemMapFs()
	gpio := mocks.NewMockGPIODriver(ctrl)
	storage := mocks.NewMockStorage(ctrl)
	installable := mocks.NewMockSubsystem(ctrl)
	cfg := aquareo.Config{}

	deviceCtrl := device.NewRPiController(fs, gpio, storage, cfg)

	t.Run("Install", func(t *testing.T) {
		installable.EXPECT().Install(deviceCtrl).Times(1).Return(nil)
		assert.NoError(t, deviceCtrl.Install(installable))
	})

	t.Run("Start", func(t *testing.T) {
		installable.EXPECT().Start().Return()
		gpio.EXPECT().Open().Return(nil)

		assert.NoError(t, deviceCtrl.Start())
	})

	t.Run("Stop", func(t *testing.T) {
		gpio.EXPECT().Close().Return(nil)
		installable.EXPECT().Stop(context.TODO())

		deviceCtrl.Stop(context.TODO())
	})
}
