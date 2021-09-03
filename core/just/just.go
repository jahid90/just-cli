package just

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/jahid90/just/core/file/json"
	"github.com/jahid90/just/core/file/text"
	"github.com/jahid90/just/core/file/yaml"
	"github.com/jahid90/just/core/just/api"
	v1 "github.com/jahid90/just/core/just/v1"
	v2 "github.com/jahid90/just/core/just/v2"
	v3 "github.com/jahid90/just/core/just/v3"
	v4 "github.com/jahid90/just/core/just/v4"
	v5 "github.com/jahid90/just/core/just/v5"
	v6 "github.com/jahid90/just/core/just/v6"
	"github.com/jahid90/just/core/logger"
)

// Version A type representing the version of a just config file
type Version struct {
	Version string `json:"version" yaml:"version"`
}

// GeneratorFn Function to generate the config
type GeneratorFn func(alias string, appendArgs []string, config interface{}) ([]*exec.Cmd, error)

func GetApi(filename string) (*api.JustApi, error) {
	ver := &Version{}

	err := parseFileContents(filename, ver)
	if err != nil {
		return nil, err
	}

	logger.Infof("found a v%s config file", ver.Version)

	switch ver.Version {
	case "1":
		just := &v1.Just{}
		err := parseFileContents(filename, just)
		if err != nil {
			return nil, err
		}
		return v1.GenerateApi(just), nil

	case "2":
		just := &v1.Just{}
		err := parseFileContents(filename, just)
		if err != nil {
			return nil, err
		}
		return v2.GenerateApi(just), nil

	case "3":
		just := &v1.Just{}
		err := parseFileContents(filename, just)
		if err != nil {
			return nil, err
		}
		return v3.GenerateApi(just), nil

	case "4":
		just := &v1.Just{}
		err := parseFileContents(filename, just)
		if err != nil {
			return nil, err
		}
		return v4.GenerateApi(just), nil

	case "5":
		just := &v5.Just{}
		err := parseFileContents(filename, just)
		if err != nil {
			return nil, err
		}
		return v5.GenerateApi(just), nil

	case "6":
		just := &v6.Just{}
		err := parseFileContents(filename, just)
		if err != nil {
			return nil, err
		}
		return v6.GenerateApi(just), nil

	default:
		return nil, errors.New("unsupported version: " + ver.Version)

	}

}

func parseFileContents(filename string, container interface{}) error {

	contents, err := text.ReadFile(filename)
	if err != nil {
		return err
	}

	logger.Debug("read the contents of the file")

	if strings.HasSuffix(filename, ".yaml") {

		logger.Debug("found a yaml config file")

		err := yaml.ParseYaml(contents, container)
		if err != nil {
			logger.Debug("parsing as yaml failed")
			return err
		}
	} else if strings.HasSuffix(filename, ".json") {

		logger.Debug("found a json config file")

		err := json.ParseJson(contents, container)
		if err != nil {
			logger.Debug("parsing as json failed")
			return err
		}

	} else {
		logger.Warnf("only json and yaml formats are supported")
		return errors.New("unknown file type")
	}

	return nil
}
