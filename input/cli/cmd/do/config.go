package do

import (
	"errors"

	"github.com/jahid90/just/core/command/executor"
	"github.com/jahid90/just/core/just"
	v1 "github.com/jahid90/just/core/just/v1"
	v2 "github.com/jahid90/just/core/just/v2"
	v3 "github.com/jahid90/just/core/just/v3"
	v4 "github.com/jahid90/just/core/just/v4"
	v5 "github.com/jahid90/just/core/just/v5"
	v6 "github.com/jahid90/just/core/just/v6"
	"github.com/urfave/cli/v2"
)

// Config Parsed config that can be used to generate and run commands
type Config struct {
	// RunCmdFunc Function to run the command corresponding to the alias received in the args
	RunCmd func(c *cli.Context) error

	// GetListingFunc Function to generate listing of available commands
	GetListing func() error

	// GetListingFunc Function to generate a short listing of available commands
	GetShortListing func() error

	// FormatFunc Function to format the config in known formats
	Format func(format string) ([]byte, error)

	// ConvertFn Function to convert configs between versions
	Convert func() ([]byte, error)
}

// Parse Parses the config file and generates a suitable Config
func Parse(contents []byte) (*Config, error) {

	c, err := just.GetConfig(contents)
	if err != nil {
		return nil, err
	}

	version := c.Version
	var cmdGeneratorFn just.GeneratorFn
	var configFile interface{}

	// we only allow the versions we know
	switch version {
	case "1":
		cmdGeneratorFn = v1.CommandGeneratorFn
		configFile = c.JustV1

	case "2":
		cmdGeneratorFn = v2.CommandGeneratorFn
		configFile = c.JustV1

	case "3":
		cmdGeneratorFn = v3.CommandGeneratorFn
		configFile = c.JustV1

	case "4":
		cmdGeneratorFn = v4.CommandGeneratorFn
		configFile = c.JustV1

	case "5":
		cmdGeneratorFn = v5.CommandGeneratorFn
		configFile = c.JustV5

	case "6":
		cmdGeneratorFn = v6.CommandGeneratorFn
		configFile = c.JustV6

	default:
		return nil, errors.New("error: unknown version: " + version)
	}

	config, err := generateConfig(c, configFile, cmdGeneratorFn)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func generateConfig(c *just.Config, configFile interface{}, fn just.GeneratorFn) (*Config, error) {
	config := &Config{
		RunCmd: func(ctx *cli.Context) error {

			alias := ctx.Args().First()
			args := ctx.Args().Tail()
			cmds, err := fn(alias, args, configFile)
			if err != nil {
				return err
			}

			err = executor.ExecuteMany(cmds)
			if err != nil {
				return err
			}

			return nil
		},
		GetListing: func() error {
			return c.ShowListing()
		},
		GetShortListing: func() error {
			return c.ShowShortListing()
		},
		Format: func(format string) ([]byte, error) {
			return c.Format(format)
		},
		Convert: func() ([]byte, error) {
			return c.Convert()
		},
	}

	return config, nil
}
