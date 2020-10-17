package justfile

import (
	"github.com/jahid90/just/lib"
)

// JustV1 A type representing v1 of just config file
type JustV1 struct {
	Version  string            `json:"version"`
	Commands map[string]string `json:"commands"`
}

// ParseV1 Parses JustV1
func ParseV1(contents []byte) (JustV1, error) {
	j := JustV1{}
	err := lib.ParseJSON(contents, &j)
	if err != nil {
		return JustV1{}, err
	}

	return j, nil
}
