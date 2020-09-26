package main

import (
	"os"

	"gopkg.in/urfave/cli.v2"
)

func main() {
	(&cli.App{}).Run(os.Args)
}
