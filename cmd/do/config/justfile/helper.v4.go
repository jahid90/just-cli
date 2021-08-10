package justfile

import (
	"os/exec"
	"strings"
)

var CommandV4GeneratorFn = func(alias string, appendArgs []string, j *Config) ([]*exec.Cmd, error) {

	entry, err := j.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	// add any additional arguments provided
	if len(appendArgs) > 0 {
		entry = entry + " " + strings.Join(appendArgs, " ")
	}

	cmd := exec.Command("sh", "-c", entry)

	return []*exec.Cmd{cmd}, nil
}
