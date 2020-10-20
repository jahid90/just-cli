package config

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/jahid90/just/cmd/do/config/justfile"
	"github.com/jahid90/just/lib/lexer"
	"github.com/jahid90/just/lib/parser"
)

var commandV3GeneratorFn = func(alias string, j *justfile.Just) (*exec.Cmd, error) {

	entry, err := j.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	fmt.Println("just @" + entry)

	lexer := lexer.NewLexer(strings.NewReader(entry))
	buffer := lexer.Run()

	buffer.Print()

	parser := parser.NewParser(buffer)
	parsed := parser.Parse()

	parsed.Print(0)

	return nil, errors.New("Warn: Not yet implemented")
}
