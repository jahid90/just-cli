package config

import (
	"errors"

	"github.com/jahid90/just/cmd/do/config/justfile"
	"github.com/jahid90/just/cmd/do/config/versioner"
	"github.com/urfave/cli/v2"
)

// RunCmdFunc Function to run the command corresponding to the alias received in the args
type RunCmdFunc func(c *cli.Context) error

// GetListingFunc Function to generate listing of available commands
type GetListingFunc func() error

// Config Parsed config that can be used to generate and run commands
type Config struct {
	RunCmd     RunCmdFunc
	GetListing GetListingFunc
}

// Fn Function to generate the config
type Fn func(j *justfile.Just) (*Config, error)

// ParseConfig Parses the config file and generates a suitable Config
func ParseConfig(contents []byte) (*Config, error) {

	version, err := versioner.ParseVersion(contents)
	if err != nil {
		return nil, err
	}

	var configFn Fn
	// we only allow the versions we know
	switch version {
	case "1":
		configFn = configFromV1
		break

	case "2":
		configFn = configFromV2
		break

	case "3":
		configFn = configFromV3
		break

	default:
		configFn = defaultFn
		break
	}

	return handle(contents, configFn)
}

func handle(contents []byte, configFn Fn) (*Config, error) {

	parserFn := justfile.GetParser()
	j, err := parserFn(contents)
	if err != nil {
		return nil, err
	}

	c, err := configFn(j)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func defaultFn(j *justfile.Just) (*Config, error) {
	return nil, errors.New("Warn: Not yet implemented")
}
