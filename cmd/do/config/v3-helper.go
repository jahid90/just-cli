package config

import (
	"fmt"

	"github.com/jahid90/just/cmd/do/config/justfile"
)

func handleV3(contents []byte) (Config, error) {
	j, err := justfile.ParseV3(contents)
	if err != nil {
		return Config{}, nil
	}

	fmt.Println("==========================")
	fmt.Println(j.Version)
	for k, v := range j.Commands {
		fmt.Println(k + ":" + v)
	}

	/*
		c, err := configFromV2(j)
		if err != nil {
			return Config{}, nil
		}

		return c, nil
	*/

	return Config{}, nil
}
