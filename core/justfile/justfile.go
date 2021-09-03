package justfile

import (
	"os/exec"

	"github.com/jahid90/just/core/file/json"
	"github.com/jahid90/just/core/file/yaml"
	v1 "github.com/jahid90/just/core/justfile/v1"
	v5 "github.com/jahid90/just/core/justfile/v5"
	v6 "github.com/jahid90/just/core/justfile/v6"
)

// Version A type representing the version of a just config file
type Version struct {
	Version string `json:"version" yaml:"version"`
}

// GeneratorFn Function to generate the config
type GeneratorFn func(alias string, appendArgs []string, config interface{}) ([]*exec.Cmd, error)

// Config A type representing a generic just config
type Config struct {
	Version            string
	JustV1             *v1.Just
	JustV5             *v5.Just
	JustV6             *v6.Just
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

// GetParserFn Returns a parser function to parse the config file
func GetConfig(contents []byte) (*Config, error) {

	v := &Version{}
	err := parseAsJsonOrYaml(contents, v)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	c.Version = v.Version
	c.JustV1 = nil
	c.JustV5 = nil
	c.JustV6 = nil

	if v.Version == "6" {

		j := &v6.Just{}
		err = parseAsJsonOrYaml(contents, j)
		if err != nil {
			return nil, err
		}

		c.JustV6 = j
		c.Format = j.Format
		c.Convert = func() ([]byte, error) { return []byte(""), nil }
		c.ShowListing = j.ShowListing
		c.ShowShortListing = j.ShowShortListing
		c.LookupAlias = j.LookupAlias
		c.LookupDependencies = j.LookupDependencies

	} else if v.Version == "5" {

		j := &v5.Just{}
		err = parseAsJsonOrYaml(contents, j)
		if err != nil {
			return nil, err
		}

		c.JustV5 = j
		c.Format = j.Format
		c.Convert = func() ([]byte, error) { return []byte(""), nil }
		c.ShowListing = j.ShowListing
		c.ShowShortListing = j.ShowShortListing
		c.LookupAlias = j.LookupAlias
		c.LookupDependencies = j.LookupDependencies

	} else {

		j := &v1.Just{}
		err = parseAsJsonOrYaml(contents, j)
		if err != nil {
			return nil, err
		}

		c.JustV1 = j
		c.Format = j.Format
		c.Convert = func() ([]byte, error) { return []byte(""), nil }
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
