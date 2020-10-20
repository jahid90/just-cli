package config

import (
	"errors"
	"os"

	"github.com/jahid90/just/cmd/do/config/justfile"
	"github.com/jahid90/just/lib/command"
	"github.com/urfave/cli/v2"
)

// Config Parsed config that can be used to generate and run commands
type Config struct {
	RunCmd     RunCmdFunc
	GetListing GetListingFunc
}

// RunCmdFunc Function to run the command corresponding to the alias received in the args
type RunCmdFunc func(c *cli.Context) error

// GetListingFunc Function to generate listing of available commands
type GetListingFunc func() error

// GeneratorFn Function to generate the config
type GeneratorFn func(j *justfile.Just) (*Config, error)

// Parse Parses the config file and generates a suitable Config
func Parse(contents []byte) (*Config, error) {

	configfileParserFn := justfile.GetParserFn()
	j, err := configfileParserFn(contents)
	if err != nil {
		return nil, err
	}

	version := j.Version
	var cmdGeneratorFn command.GeneratorFn

	// we only allow the versions we know
	switch version {
	case "1":
		cmdGeneratorFn = commandV1GeneratorFn
		break

	case "2":
		cmdGeneratorFn = commandV2GeneratorFn
		break

	case "3":
		cmdGeneratorFn = commandV3GeneratorFn
		break

	default:
		return nil, errors.New("Error: unknown version: " + version)
	}

	config, err := generateConfig(j, cmdGeneratorFn)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func generateConfig(j *justfile.Just, fn command.GeneratorFn) (*Config, error) {
	c := &Config{
		RunCmd: func(c *cli.Context) error {

			alias := c.Args().First()
			cmd, err := fn(alias, j)
			if err != nil {
				return err
			}

			err = command.Run(cmd, os.Stdout, os.Stderr)
			if err != nil {
				return err
			}

			return nil
		},
		GetListing: j.ShowListing,
	}

	return c, nil
}
