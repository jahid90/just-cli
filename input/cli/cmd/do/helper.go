package do

import (
	"github.com/jahid90/just/core/file/text"
	"github.com/urfave/cli/v2"
)

func parseConfig(ctx *cli.Context) (*Config, error) {

	contents, err := readConfigFile(ctx)
	if err != nil {
		return nil, err
	}

	c, err := Parse(contents)
	if err != nil {
		return nil, err
	}

	return c, nil

}

func readConfigFile(ctx *cli.Context) ([]byte, error) {

	var configFileName = ctx.String("config-file")

	if len(configFileName) != 0 {

		contents, err := text.ReadFile(configFileName)
		if err != nil {
			return nil, err
		}

		return contents, nil

	} else {

		contents, err := text.ReadFile("just.json")
		if err != nil {

			contents, err := text.ReadFile("just.yaml")
			if err != nil {
				return nil, err
			}

			return contents, nil
		}

		return contents, nil

	}
}
