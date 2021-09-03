package main

import (
	"github.com/jahid90/just/core/config"
	"github.com/jahid90/just/core/logger"
	"github.com/jahid90/just/input/cli"
	"github.com/jahid90/just/output/console/colorize"
	"github.com/jahid90/just/output/console/plain"
)

var Environment string
var GitCommit string

func main() {

	// inject implementations to injection points
	logger.Logger = plain.Println
	logger.Formatter = plain.Sprintf
	logger.Colorizer = colorize.Sprint

	config.SetLogLevel(Environment)

	if len(GitCommit) == 0 {
		GitCommit = "commit unknown"
	}

	cli.Run(GitCommit)
}
