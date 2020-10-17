package versioner

import (
	"errors"

	"github.com/jahid90/just/lib"
)

type versioner struct {
	Version string `json:"version"`
}

// ParseVersion Parses version from config
func ParseVersion(config []byte) (string, error) {

	v := versioner{}
	err := lib.ParseJSON(config, &v)
	if err != nil {
		return "", err
	}

	if v.Version == "" {
		return "", errors.New("Error: config file is missing version")
	}

	return v.Version, nil
}
