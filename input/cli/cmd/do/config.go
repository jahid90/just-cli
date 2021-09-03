package do

import (
	"errors"
	"os/exec"

	"github.com/jahid90/just/core/command/executor"
	"github.com/jahid90/just/core/justfile"
	v1 "github.com/jahid90/just/core/justfile/v1"
	v2 "github.com/jahid90/just/core/justfile/v2"
	v3 "github.com/jahid90/just/core/justfile/v3"
	v4 "github.com/jahid90/just/core/justfile/v4"
	v5 "github.com/jahid90/just/core/justfile/v5"
	v6 "github.com/jahid90/just/core/justfile/v6"
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
type GeneratorFn func(alias string, appendArgs []string, config interface{}) ([]*exec.Cmd, error)

// ConvertFn Function to convert configs between versions
type ConvertFn func() ([]byte, error)

// Parse Parses the config file and generates a suitable Config
func Parse(contents []byte) (*Config, error) {

	c, err := justfile.GetConfig(contents)
	if err != nil {
		return nil, err
	}

	version := c.Version
	var cmdGeneratorFn GeneratorFn
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

func generateConfig(c *justfile.Config, configFile interface{}, fn GeneratorFn) (*Config, error) {
	config := &Config{
		RunCmd: func(ctx *cli.Context) error {

			alias := ctx.Args().First()
			cmds, err := fn(alias, ctx.Args().Tail(), configFile)
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
