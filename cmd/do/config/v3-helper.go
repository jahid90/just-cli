package config

import (
	"fmt"
	"strings"

	"github.com/jahid90/just/lib/lexer"
	"github.com/jahid90/just/lib/parser"

	"github.com/jahid90/just/cmd/do/config/justfile"
	"github.com/urfave/cli/v2"
)

func configFromV3(j *justfile.Just) (*Config, error) {
	c := &Config{
		RunCmd: func(c *cli.Context) error {

			for alias, command := range j.Commands {
				fmt.Println("just @" + alias)

				lexer := lexer.NewLexer(strings.NewReader(command))
				buffer := lexer.Run()

				buffer.Print()

				parser := parser.NewParser(buffer)
				parsed := parser.Parse()

				parsed.Print(0)
			}

			return nil
		},
		GetListing: j.ShowListing,
	}

	return c, nil
}
