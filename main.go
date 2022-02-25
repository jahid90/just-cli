package main

import (
	"os"

	"github.com/jahid90/just/core/config"
	"github.com/jahid90/just/core/logger"
	"github.com/jahid90/just/input/cli"
	"github.com/jahid90/just/output/console"
)

var Environment string = "development"
var GitCommit string = "devel  version"

func main() {

	plainConsole := console.NewPlainConsole(os.Stderr)
	coloredConsole := console.NewColoredConsole(os.Stderr)

	// inject implementations to injection points
	logger.Logger = plainConsole.Println
	logger.Formatter = plainConsole.Sprintf
	logger.Colorizer = coloredConsole.Sprint

	config.SetLogLevel(Environment)

	cli.Run(GitCommit)
}
