package yaml

import (
	"errors"

	"github.com/jahid90/just/core/file/text"
	"gopkg.in/yaml.v2"
)

// ParseJSON Parses yaml data into the container
func ParseYaml(data []byte, container interface{}) error {
	err := yaml.Unmarshal(data, container)
	if err != nil {
		return errors.New("error: bad config file format: " + err.Error())
	}

	return nil
}

// ParseYamlFromFile Parses a yaml file into the container
func ParseYamlFromFile(filename string, container interface{}) error {
	contents, err := text.ReadFile(filename)
	if err != nil {
		return err
	}

	return ParseYaml(contents, container)
}
