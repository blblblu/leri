package commands

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"

	"github.com/blblblu/leri/lib/parsing"
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
		Usage: "generate markdown documentation for a source file",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "input",
				Usage:       "the source input file `path` to use",
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

			parser, err := getParserForExtension(args.input)
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed getting parser: %s", err.Error()), 2)
			}

			inputData, err := ioutil.ReadFile(args.input)
			if err != nil {
				return cli.Exit(fmt.Sprintf("failed reading input file: %s", err.Error()), 3)
			}

			mdString := parser.GenerateMd(string(inputData))

			if err := ioutil.WriteFile(args.output, []byte(mdString), 0644); err != nil {
				return cli.Exit(fmt.Sprintf("failed writing output file: %s", err.Error()), 4)
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

func getParserForExtension(input string) (parsing.Parser, error) {
	ext := path.Ext(input)
	switch ext {
	case ".go":
		p := parsing.NewSimpleGoParser()
		return &p, nil
	case ".sas":
		p := parsing.NewSimpleSasParser()
		return &p, nil
	}
	return nil, fmt.Errorf("input file type \"%s\" not supported", ext)
}
