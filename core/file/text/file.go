package text

import (
	"errors"
	"io/ioutil"
)

// ReadFile Reads filename and returns its contents
func ReadFile(filename string) ([]byte, error) {
	config, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("error: config file not found")
	}

	return config, nil
}
