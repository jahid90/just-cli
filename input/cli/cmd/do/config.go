package do

import (
	"errors"

	"github.com/jahid90/just/core/file/text"
	"github.com/jahid90/just/core/logger"
	"github.com/urfave/cli/v2"
)

func getConfigFileName(ctx *cli.Context) (string, error) {
	var configFileName string

	configFileName = ctx.String("config-file")
	if len(configFileName) != 0 {
		return configFileName, nil
	}

	logger.Debug("no config-file flag provided; falling back to just.yaml")

	configFileName = "just.yaml"
	if text.Exists(configFileName) {
		return configFileName, nil
	}

	logger.Debug("no just.yaml file found; falling back to just.json")

	configFileName = "just.json"

	if text.Exists(configFileName) {
		return configFileName, nil
	}

	return "", errors.New("no config file was found")
}
