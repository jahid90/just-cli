package config

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/jahid90/just/cmd/do/config/justfile"
)

var commandV4GeneratorFn = func(alias string, appendArgs []string, j *justfile.Just) (*exec.Cmd, error) {

	entry, err := j.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	// add any additional arguments provided
	if len(appendArgs) > 0 {
		entry = entry + " " + strings.Join(appendArgs, " ")
	}

	fmt.Println("just @" + entry)

	cmd := exec.Command("sh", "-c", entry)

	return cmd, nil
}
