package config

import (
	"errors"

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

// ParseConfig Parses the config file and generates a suitable Config
func ParseConfig(contents []byte) (Config, error) {

	version, err := versioner.ParseVersion(contents)
	if err != nil {
		return Config{}, err
	}

	// we only allow the versions we know
	switch version {
	case "1":
		config, err := handleV1(contents)
		if err != nil {
			return Config{}, err
		}
		return config, nil

	case "2":
		config, err := handleV2(contents)
		if err != nil {
			return Config{}, err
		}
		return config, nil

	case "3":
		_, err := handleV3(contents)
		if err != nil {
			return Config{}, nil
		}
		return Config{}, errors.New("Warn: Not yet implemented")

	default:
		return Config{}, errors.New("Error: unknown version: " + version)
	}
}
