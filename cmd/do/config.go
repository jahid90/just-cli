package do

import "errors"

var configFileName = "just.json"

// TODO - need a way to return generic versions, not just justV1
func parseConfig() (justV1, error) {

	version, err := parseConfigVersion()
	if err != nil {
		return justV1{}, err
	}

	// we only allow the versions we know
	switch version {
	case "1":
		config, err := parseConfigV1()
		if err != nil {
			return justV1{}, nil
		}
		return config, nil
	default:
		return justV1{}, errors.New("Error: unknown version: " + version)
	}
}
