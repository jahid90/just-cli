package do

import (
	"github.com/jahid90/just/cmd/do/config"
	"github.com/jahid90/just/lib"
	"github.com/urfave/cli/v2"
)

func parseConfig(ctx *cli.Context) (*config.Config, error) {

	contents, err := readConfigFile(ctx)
	if err != nil {
		return nil, err
	}

	c, err := config.Parse(contents)
	if err != nil {
		return nil, err
	}

	return c, nil

}

func readConfigFile(ctx *cli.Context) ([]byte, error) {

	var configFileName = ctx.String("config-file")

	if len(configFileName) != 0 {

		contents, err := lib.ReadFile(configFileName)
		if err != nil {
			return nil, err
		}

		return contents, nil

	} else {

		contents, err := lib.ReadFile("just.json")
		if err != nil {

			contents, err := lib.ReadFile("just.yaml")
			if err != nil {
				return nil, err
			}

			return contents, nil
		}

		return contents, nil

	}
}
