package do

import (
	"github.com/jahid90/just/cmd/do/config"
	"github.com/jahid90/just/lib"
	"github.com/urfave/cli/v2"
)

func parseConfig(ctx *cli.Context) (config.Config, error) {

	var configFileName = ctx.String("config-file")

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
