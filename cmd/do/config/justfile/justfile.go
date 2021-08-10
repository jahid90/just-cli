package justfile

import (
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
}

type ShowListingFn func() error
type LookupAliasFn func(string) (string, error)

// ParserFn A function representing a parser
// On being invoked, parses the contents passed to it, generates a Just config file and returns a pointer to it
type ParserFn func([]byte) (*Config, error)

// GetParserFn Returns a parser function to parse the config file
func GetConfig(contents []byte) (*Config, error) {

	v := &Version{}
	err := lib.ParseJSON(contents, v)
	if err != nil {
		err := lib.ParseYaml(contents, v)
		if err != nil {
			return nil, err
		}
	}

	if v.Version == "5" {

		j := &JustV5{}

		c := &Config{}
		c.Version = v.Version
		c.just = nil
		c.justV5 = j
		c.ShowListing = j.ShowListing
		c.LookupAlias = j.LookupAlias

		return c, nil
	}

	j := &Just{}
	err = lib.ParseJSON(contents, j)
	if err != nil {
		err := lib.ParseYaml(contents, j)
		if err != nil {
			return nil, err
		}
	}

	c := &Config{}
	c.Version = v.Version
	c.just = j
	c.justV5 = nil
	c.ShowListing = j.ShowListing
	c.LookupAlias = j.LookupAlias

	return c, nil
}
