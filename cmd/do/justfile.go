package do

import (
	"encoding/json"
	"errors"
	"io/ioutil"
)

var configFileName = "just.json"

type versioner struct {
	Version string `json:"version"`
}

type justV1 struct {
	Version  string            `json:"version"`
	Commands map[string]string `json:"commands"`
}

func readFile() ([]byte, error) {
	config, err := ioutil.ReadFile(configFileName)
	if err != nil {
		return nil, errors.New("config file not found")
	}

	return config, nil
}

func parseJSON(data []byte, container interface{}) error {
	err := json.Unmarshal(data, container)
	if err != nil {
		return errors.New("bad config file format: " + err.Error())
	}

	return nil
}

func parseConfigVersion() (string, error) {

	json, err := readFile()
	if err != nil {
		return "", err
	}

	v := versioner{}
	err = parseJSON(json, &v)
	if err != nil {
		return "", err
	}

	if v.Version == "" {
		return "", errors.New("config file is missing version")
	}

	return v.Version, nil
}

func parseConfigV1() (justV1, error) {
	json, err := readFile()
	if err != nil {
		return justV1{}, err
	}

	config := justV1{}
	err = parseJSON(json, &config)
	if err != nil {
		return justV1{}, nil
	}

	return config, nil
}

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
		return justV1{}, errors.New("unknown version: " + version)
	}
}
