package text

import (
	"errors"
	"io/ioutil"
	"os"

	"github.com/jahid90/just/core/logger"
)

// ReadFile Reads filename and returns its contents
func ReadFile(filename string) ([]byte, error) {
	contents, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, errors.New("config file not found")
	}

	return contents, nil
}

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}

	if os.IsNotExist(err) {
		logger.Info("file does not exist")
	} else {
		logger.Error(err.Error())
	}

	return false
}
