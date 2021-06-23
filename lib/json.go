package lib

import (
	"encoding/json"
	"errors"
)

// ParseJSON Parses data as json into container
func ParseJSON(data []byte, container interface{}) error {
	err := json.Unmarshal(data, container)
	if err != nil {
		return errors.New("error: bad config file format: " + err.Error())
	}

	return nil
}
