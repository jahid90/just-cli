package v3

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/jahid90/just/core/just/api"
	v1 "github.com/jahid90/just/core/just/v1"
	"github.com/jahid90/just/core/logger"
	"github.com/jahid90/just/core/misc/lexer"
	"github.com/jahid90/just/core/misc/parser"
)

var CommandGeneratorFn = func(alias string, appendArgs []string, config interface{}) ([]*exec.Cmd, error) {

	j, ok := config.(*v1.Just)
	if !ok {
		logger.Error("bad type; expected v1 Just")
		return nil, errors.New("internal error")
	}

	aka, err := j.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	entry, ok := aka.(string)
	if !ok {
		return nil, errors.New("error: internal - unexpected type received")
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

func GenerateApi(config *v1.Just) *api.JustApi {
	return v1.GenerateApi(config)
}
