package common

import "github.com/urfave/cli/v2"

func GetNameFlag(usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   usage,
	}
}

func GetPathFlag(usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "path",
		Aliases: []string{"p"},
		Usage:   usage,
	}
}

func GetAttributeFlag(usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "attribute",
		Aliases: []string{"a"},
		Usage:   usage,
	}
}
