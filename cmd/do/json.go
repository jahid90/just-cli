package do

import (
	"encoding/json"
	"errors"
)

func parseJSON(data []byte, container interface{}) error {
	err := json.Unmarshal(data, container)
	if err != nil {
		return errors.New("Error: bad config file format: " + err.Error())
	}

	return nil
}
