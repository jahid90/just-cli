package justfile

import (
	"os/exec"
	"strings"

	"github.com/jahid90/just/lib/command"
)

var CommandV1GeneratorFn = func(alias string, appendArgs []string, j *Config) ([]*exec.Cmd, error) {

	entry, err := j.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	// add any additional arguments provided
	if len(appendArgs) > 0 {
		entry = entry + " " + strings.Join(appendArgs, " ")
	}

	commandLine := strings.Split(entry, " ")
	c := commandLine[0]
	args := commandLine[1:]

	err = command.Validate(c)
	if err != nil {
		return nil, err
	}

	// generate the command; ignore any additional arguments supplied
	cmd := exec.Command(c, args...)

	return []*exec.Cmd{cmd}, nil
}
