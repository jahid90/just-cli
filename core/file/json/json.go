package json

import (
	"encoding/json"
	"errors"
)

// ParseJson Parses data as json into container
func ParseJson(data []byte, container interface{}) error {
	err := json.Unmarshal(data, container)
	if err != nil {
		return errors.New("error: bad config file format: " + err.Error())
	}

	return nil
}
