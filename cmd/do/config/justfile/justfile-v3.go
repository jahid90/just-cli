package justfile

import (
	"fmt"
	"strings"

	"github.com/jahid90/just/lib"
)

// JustV3 A type representing v3 of the config file
type JustV3 struct {
	Version        string            `json:"version"`
	Commands       map[string]string `json:"commands"`
	ParsedCommands map[string]CommandV3
}

// CommandV3 Represents a command in a v3 config file
type CommandV3 struct {
}

// ParseV3 parses JustV3
func ParseV3(contents []byte) (JustV3, error) {
	j := JustV3{}
	err := lib.ParseJSON(contents, &j)
	if err != nil {
		return JustV3{}, err
	}

	for alias, command := range j.Commands {
		fmt.Println()
		fmt.Println("Processing => " + alias + ": " + command)
		fmt.Println("-------------------------------------------------------------------------------------------")
		lexer := lib.NewLexer(strings.NewReader(command))

		for {
			token := lexer.Lex()

			if token.Type == lib.EOF {
				break
			}

			fmt.Println(token)
		}
	}

	return j, nil
}
