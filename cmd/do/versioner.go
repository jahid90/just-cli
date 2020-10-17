package do

import "errors"

type versioner struct {
	Version string `json:"version"`
}

func parseConfigVersion() (string, error) {

	json, err := readFile(configFileName)
	if err != nil {
		return "", err
	}

	v := versioner{}
	err = parseJSON(json, &v)
	if err != nil {
		return "", err
	}

	if v.Version == "" {
		return "", errors.New("Error: config file is missing version")
	}

	return v.Version, nil
}
