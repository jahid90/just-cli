package config

import (
	"errors"

	"github.com/jahid90/just/core/command/executor"
	"github.com/jahid90/just/input/cli/cmd/do/config/justfile"
	"github.com/urfave/cli/v2"
)

// Config Parsed config that can be used to generate and run commands
type Config struct {
	RunCmd          RunCmdFunc
	GetListing      GetListingFunc
	GetShortListing GetShortListingFunc
	Format          FormatFunc
	Convert         ConvertFn
}

// RunCmdFunc Function to run the command corresponding to the alias received in the args
type RunCmdFunc func(c *cli.Context) error

// GetListingFunc Function to generate listing of available commands
type GetListingFunc func() error

// GetListingFunc Function to generate a short listing of available commands
type GetShortListingFunc func() error

// FormatFunc Function to format the config in known formats
type FormatFunc func(format string) ([]byte, error)

// GeneratorFn Function to generate the config
type GeneratorFn func(c *justfile.Config) (*Config, error)

// ConvertFn Function to convert configs between versions
type ConvertFn func() ([]byte, error)

// Parse Parses the config file and generates a suitable Config
func Parse(contents []byte) (*Config, error) {

	c, err := justfile.GetConfig(contents)
	if err != nil {
		return nil, err
	}

	version := c.Version
	var cmdGeneratorFn justfile.GeneratorFn

	// we only allow the versions we know
	switch version {
	case "1":
		cmdGeneratorFn = justfile.CommandV1GeneratorFn

	case "2":
		cmdGeneratorFn = justfile.CommandV2GeneratorFn

	case "3":
		cmdGeneratorFn = justfile.CommandV3GeneratorFn

	case "4":
		cmdGeneratorFn = justfile.CommandV4GeneratorFn

	case "5":
		cmdGeneratorFn = justfile.CommandV5GeneratorFn

	case "6":
		cmdGeneratorFn = justfile.CommandV6GeneratorFn

	default:
		return nil, errors.New("error: unknown version: " + version)
	}

	config, err := generateConfig(c, cmdGeneratorFn)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func generateConfig(c *justfile.Config, fn justfile.GeneratorFn) (*Config, error) {
	config := &Config{
		RunCmd: func(ctx *cli.Context) error {

			alias := ctx.Args().First()
			cmd, err := fn(alias, ctx.Args().Tail(), c)
			if err != nil {
				return err
			}

			err = executor.Run(cmd)
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
