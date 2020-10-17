package do

var configFileName = "just.json"

type justV1 struct {
	Version  string            `json:"version"`
	Commands map[string]string `json:"commands"`
}

func parseConfigV1() (justV1, error) {
	json, err := readFile(configFileName)
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
