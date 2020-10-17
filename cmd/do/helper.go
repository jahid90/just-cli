package do

import (
	"github.com/jahid90/just/cmd/do/config"
	"github.com/jahid90/just/lib"
)

var configFileName = "just.json"

func parseConfig() (config.Config, error) {

	contents, err := lib.ReadFile(configFileName)
	if err != nil {
		return config.Config{}, err
	}

	c, err := config.ParseConfig(contents)
	if err != nil {
		return config.Config{}, err
	}

	return c, nil
}
