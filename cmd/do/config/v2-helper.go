package config

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/jahid90/just/cmd/do/config/justfile"
	"github.com/jahid90/just/lib/command"
)

var commandV2GeneratorFn = func(alias string, j *justfile.Just) (*exec.Cmd, error) {

	entry, err := j.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	// output the command we are running
	fmt.Println("just @" + entry)

	commandLine := strings.Split(entry, " ")
	c := commandLine[0]
	args := commandLine[1:]

	err = command.Validate(c)
	if err != nil {
		return nil, err
	}

	// generate the command; ignore any additional arguments supplied
	cmd := exec.Command(c, args...)

	return cmd, nil
}
