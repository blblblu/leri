package main

import (
	"os"

	"github.com/blblblu/leri/lib/commands"
	"github.com/urfave/cli"
)

var (
	version = "master"
)

func main() {
	app := cli.App{
		Name:  "leri",
		Usage: "a markdown documentation genarator for source code files",
		Authors: []*cli.Author{
			{Name: "Sebastian Schulz", Email: "mail@sesc.me"},
		},
		Version: version,
		Commands: []*cli.Command{
			commands.NewGenMdCommand(),
		},
	}

	app.Run(os.Args)
}
