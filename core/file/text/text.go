package text

import (
	"errors"
	"io/ioutil"
)

// ReadFile Reads filename and returns its contents
func ReadFile(filename string) ([]byte, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("config file not found")
	}

	return contents, nil
}
