package daemon_test

import (
	"context"
	"os"
	"testing"

	"github.com/fnandes/aquareo/internal/daemon"
	"github.com/spf13/afero"
	"go.etcd.io/bbolt"
)

func TestStart(t *testing.T) {
	fs := afero.NewMemMapFs()
	cfgFile, err := fs.Create("config.json")
	if err != nil {
		panic(err)
	}

	if _, err := cfgFile.WriteString("{}"); err != nil {
		panic(err)
	}

	db := MustCreateTempDbFile()
	defer os.Remove(db.Path())

	d := daemon.NewDaemon(fs, db)

	d.Start()
	defer d.Stop(context.TODO())
}

func MustCreateTempDbFile() *bbolt.DB {
	file, err := os.CreateTemp("", "aq-test-daemon-")
	if err != nil {
		panic(err)
	}

	db, err := bbolt.Open(file.Name(), 0600, nil)
	if err != nil {
		panic(err)
	}

	return db
}
