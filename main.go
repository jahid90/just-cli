package main

import (
	"github.com/jahid90/just/core/logger"
	"github.com/jahid90/just/input/cli"
	"github.com/jahid90/just/output/console/colorize"
	"github.com/jahid90/just/output/console/plain"
)

func main() {

	// inject implementations to injection points
	logger.Logger = plain.Println
	logger.Formatter = plain.Sprintf
	logger.Colorizer = colorize.Sprint

	cli.Run()
}
