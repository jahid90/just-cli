package cmd

import (
	"github.com/jahid90/just/cmd/do"
	"github.com/jahid90/just/cmd/hello"
	"github.com/urfave/cli/v2"
)

// GetSubCommands returns the subcommands for the app.
func GetSubCommands() []*cli.Command {

	doCmd := do.Cmd()
	helloCmd := hello.Cmd()

	return []*cli.Command{
		doCmd,
		helloCmd,
	}
}
