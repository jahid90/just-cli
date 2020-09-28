package cmd

import (
	"github.com/jahid90/just/cmd/hello"
	"github.com/urfave/cli/v2"
)

// GetSubCommands returns the subcommands for the app.
func GetSubCommands() []*cli.Command {

	helloCmd := hello.Cmd()

	return []*cli.Command{
		helloCmd,
	}
}
