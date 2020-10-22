package config

import (
	"fmt"
	"os/exec"

	"github.com/jahid90/just/cmd/do/config/justfile"
)

var commandV4GeneratorFn = func(alias string, j *justfile.Just) (*exec.Cmd, error) {

	entry, err := j.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	fmt.Println("just @" + entry)

	cmd := exec.Command("sh", "-c", entry)

	return cmd, nil
}
