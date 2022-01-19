package settings

import (
	_ "embed"

	"gopkg.in/yaml.v2"
)

//go:embed conf.yml
var confFile []byte

func New() (*Settings, error) {
	config := &Settings{}

	err := yaml.Unmarshal(confFile, &config)

	return config, err
}
