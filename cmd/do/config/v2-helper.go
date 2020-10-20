package config

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/jahid90/just/cmd/do/config/justfile"
)

var commandV2GeneratorFn = func(alias string, j *justfile.Just) (*exec.Cmd, error) {

	// check if the alias is present in the config file
	entry, ok := j.Commands[alias]
	if !ok {
		return nil, errors.New("Error: alias `" + alias + "` not found in the config file")
	}

	commandLine := strings.Split(entry, " ")
	command := commandLine[0]
	args := commandLine[1:]

	// check that the command exists
	_, err := exec.LookPath(command)
	if err != nil {
		return nil, errors.New("Error: " + command + " - command not found")
	}

	// output the command we are running
	fmt.Println("just @" + entry)

	// generate the command; ignore any additional arguments supplied
	cmd := exec.Command(command, args...)

	return cmd, nil
}
