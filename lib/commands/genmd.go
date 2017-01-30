package commands

import (
	"fmt"
	"os"
	"strings"

	"github.com/urfave/cli"
)

type genArgs struct {
	input  string
	output string
}

func NewGenMdCommand() *cli.Command {
	args := genArgs{}

	return &cli.Command{
		Name:  "gen",
		Usage: "generate markdown documentation for a sas file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "input",
				Usage:       "the SAS input file `path` to use",
				Aliases:     []string{"i"},
				Destination: &args.input,
			},
			&cli.StringFlag{
				Name:        "output",
				Usage:       "the markdown output file `path` to use",
				Aliases:     []string{"o"},
				Destination: &args.output,
			},
		},
		Action: func(ctx *cli.Context) error {
			if err := checkGenArgs(ctx, &args); err != nil {
				fmt.Fprintf(os.Stderr, "%s\n\n", err)
				cli.ShowCommandHelp(ctx, "gen")
				return cli.Exit("", 1)
			}

			return nil
		},
	}
}

func checkGenArgs(ctx *cli.Context, args *genArgs) error {
	errors := []string{}
	if args.input == "" {
		errors = append(errors, "input file path must be set")
	}
	if args.output == "" {
		errors = append(errors, "output file path must be set")
	}

	if len(errors) > 0 {
		errorMessage := strings.Join(errors, "\n")
		return fmt.Errorf(errorMessage)
	}

	return nil
}
