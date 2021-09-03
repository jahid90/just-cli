package justfile

import (
	"os/exec"

	"github.com/jahid90/just/core/file/json"
	"github.com/jahid90/just/core/file/yaml"
	v6 "github.com/jahid90/just/input/cli/cmd/do/config/justfile/v6"
)

// Version A type representing the version of a just config file
// Is used to determine the appropriate container to unmarshall the config file into
type Version struct {
	Version string `json:"version" yaml:"version"`
}

// Config A type representing a generic just config
type Config struct {
	just               *Just
	justV5             *JustV5
	justV6             *v6.Just
	Version            string
	Format             FormatFn
	Convert            ConvertFn
	ShowListing        ShowListingFn
	ShowShortListing   ShowShortListingFn
	LookupAlias        LookupAliasFn
	LookupDependencies LookupDependenciesFn
}

type ConvertFn func() ([]byte, error)
type FormatFn func(format string) ([]byte, error)
type ShowListingFn func() error
type ShowShortListingFn func() error
type LookupAliasFn func(alias string) (interface{}, error)
type LookupDependenciesFn func(alias string) ([]string, error)

// GeneratorFn A function to generate an exec.Cmd that can be run
type GeneratorFn func(string, []string, *Config) ([]*exec.Cmd, error)

// GetParserFn Returns a parser function to parse the config file
func GetConfig(contents []byte) (*Config, error) {

	v := &Version{}
	err := parseAsJsonOrYaml(contents, v)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	c.Version = v.Version
	c.just = nil
	c.justV5 = nil
	c.justV6 = nil

	if v.Version == "6" {

		j := &v6.Just{}
		err = parseAsJsonOrYaml(contents, j)
		if err != nil {
			return nil, err
		}

		c.justV6 = j
		c.Format = j.Format
		c.Convert = j.Convert
		c.ShowListing = j.ShowListing
		c.ShowShortListing = j.ShowShortListing
		c.LookupAlias = j.LookupAlias
		c.LookupDependencies = j.LookupDependencies

	} else if v.Version == "5" {

		j := &JustV5{}
		err = parseAsJsonOrYaml(contents, j)
		if err != nil {
			return nil, err
		}

		c.justV5 = j
		c.Format = j.Format
		c.Convert = j.Convert
		c.ShowListing = j.ShowListing
		c.ShowShortListing = j.ShowShortListing
		c.LookupAlias = j.LookupAlias
		c.LookupDependencies = j.LookupDependencies

	} else {

		j := &Just{}
		err = parseAsJsonOrYaml(contents, j)
		if err != nil {
			return nil, err
		}

		c.just = j
		c.Format = j.Format
		c.Convert = j.Convert
		c.ShowListing = j.ShowListing
		c.ShowShortListing = j.ShowShortListing
		c.LookupAlias = j.LookupAlias
		c.LookupDependencies = j.LookupDependencies

	}

	return c, nil
}

func parseAsJsonOrYaml(contents []byte, container interface{}) error {

	err := json.ParseJson(contents, container)
	if err != nil {
		err := yaml.ParseYaml(contents, container)
		if err != nil {
			return err
		}
	}

	return nil
}
