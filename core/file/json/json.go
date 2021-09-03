package json

import (
	"encoding/json"
	"errors"

	"github.com/jahid90/just/core/file/text"
)

// ParseJson Parses json data into the container
func ParseJson(data []byte, container interface{}) error {
	err := json.Unmarshal(data, container)
	if err != nil {
		return errors.New("error: bad config file format: " + err.Error())
	}

	return nil
}

// ParseJsonFromFile Parses a json file into the container
func ParseJsonFromFile(filename string, container interface{}) error {
	contents, err := text.ReadFile(filename)
	if err != nil {
		return err
	}

	return ParseJson(contents, container)
}
