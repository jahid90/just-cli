package justfile

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/jahid90/just/lib/lexer"
	"github.com/jahid90/just/lib/parser"
)

var CommandV3GeneratorFn = func(alias string, appendArgs []string, j *Config) ([]*exec.Cmd, error) {

	entry, err := j.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	// add any additional arguments provided
	if len(appendArgs) > 0 {
		entry = entry + " " + strings.Join(appendArgs, " ")
	}

	lexer := lexer.NewLexer(strings.NewReader(entry))
	buffer := lexer.Run()

	buffer.Print()

	parser := parser.NewParser(buffer)
	parsed := parser.Parse()

	parsed.Print(0)

	return nil, errors.New("warn: not yet implemented")
}
