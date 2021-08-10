package justfile

import (
	"os/exec"

	"github.com/jahid90/just/lib"
)

// Version A type representing the version of a just config file
// Is used to determine the appropriate container to unmarshall the config file into
type Version struct {
	Version string `json:"version" yaml:"version"`
}

// Config A type representing a generic just config
type Config struct {
	just        *Just
	justV5      *JustV5
	Version     string
	ShowListing ShowListingFn
	LookupAlias LookupAliasFn
	Format      FormatFn
}

type FormatFn func(string) ([]byte, error)
type ShowListingFn func() error
type LookupAliasFn func(string) (string, error)

// GeneratorFn A function to generate an exec.Cmd that can be run
type GeneratorFn func(string, []string, *Config) (*exec.Cmd, error)

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

	if v.Version == "5" {

		j := &JustV5{}
		err = parseAsJsonOrYaml(contents, j)
		if err != nil {
			return nil, err
		}

		c.justV5 = j
		c.Format = j.Format
		c.ShowListing = j.ShowListing
		c.LookupAlias = j.LookupAlias

	} else {

		j := &Just{}
		err = parseAsJsonOrYaml(contents, j)
		if err != nil {
			return nil, err
		}

		c.just = j
		c.Format = j.Format
		c.ShowListing = j.ShowListing
		c.LookupAlias = j.LookupAlias

	}

	return c, nil
}

func parseAsJsonOrYaml(contents []byte, container interface{}) error {

	err := lib.ParseJson(contents, container)
	if err != nil {
		err := lib.ParseYaml(contents, container)
		if err != nil {
			return err
		}
	}

	return nil
}
