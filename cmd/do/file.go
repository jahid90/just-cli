package do

import (
	"errors"
	"io/ioutil"
)

func readFile(filename string) ([]byte, error) {
	config, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("Error: config file not found")
	}

	return config, nil
}
