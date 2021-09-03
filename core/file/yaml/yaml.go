package yaml

import (
	"errors"

	"gopkg.in/yaml.v2"
)

// ParseJSON Parses data as json into container
func ParseYaml(data []byte, container interface{}) error {
	err := yaml.Unmarshal(data, container)
	if err != nil {
		return errors.New("error: bad config file format: " + err.Error())
	}

	return nil
}
