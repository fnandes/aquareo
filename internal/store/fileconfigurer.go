package store

import (
	"encoding/json"
	"fmt"

	"github.com/pedrobfernandes/aquareo/internal/aquareo"
	"github.com/spf13/afero"
)

type fileConf struct {
	name     string
	conf     aquareo.Config
	fs       afero.Fs
	isLoaded bool
}

func NewFileConfigurer(name string, fs afero.Fs) *fileConf {
	return &fileConf{name: name, isLoaded: false}
}

func (c *fileConf) Get() (aquareo.Config, error) {
	if !c.isLoaded {
		buf, err := afero.ReadFile(c.fs, c.name)
		if err != nil {
			return aquareo.Config{}, fmt.Errorf("Get: failed to load config file: %v", err)
		}
		if err := json.Unmarshal(buf, &c.conf); err != nil {
			return aquareo.Config{}, fmt.Errorf("Get: failed to parse config file: %v", err)
		}
	}
	return c.conf, nil
}

func (l *fileConf) Save(cfg aquareo.Config) error {
	return nil
}
