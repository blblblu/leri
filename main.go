package main

import (
	"os"

	"github.com/blblblu/sasdoc/lib/commands"
	"github.com/urfave/cli"
)

var (
	version = "master"
)

func main() {
	app := cli.App{
		Name:    "SASdoc",
		Usage:   "a markdown documentation genarator for SAS files",
		Version: version,
		Commands: []*cli.Command{
			commands.NewGenMdCommand(),
		},
	}

	app.Run(os.Args)
}
