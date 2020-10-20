package justfile

import (
	"github.com/jahid90/just/lib"
)

// JustV2 A type representing v2 of just config file; same as JustV1
type JustV2 struct {
	Version  string            `json:"version"`
	Commands map[string]string `json:"commands"`
}

// ParseV2 parses JustV2
func ParseV2(contents []byte) (JustV2, error) {
	j := JustV2{}
	err := lib.ParseJSON(contents, &j)
	if err != nil {
		return JustV2{}, err
	}

	return j, nil
}
