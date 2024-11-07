package common

import (
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func GetNameFlag(usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "name",
		Aliases: []string{"n"},
		Usage:   usage,
	}
}

func GetNameFlagValue(cCtx *cli.Context) (string, error) {
	name := cCtx.String("name")

	if cCtx.IsSet("name") && name == "" {
		return "", fmt.Errorf("-n flag requires a string value")
	}

	return name, nil
}

func GetPathFlag(usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "path",
		Aliases: []string{"p"},
		Usage:   usage,
	}
}

func GetPathFlagValue(cCtx *cli.Context) (string, error) {
	path := cCtx.String("path")

	if cCtx.IsSet("path") && path == "" {
		return "", fmt.Errorf("-p flag requires a string value")
	}

	if path == "" {
		path = "."
	}

	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("could not determine absolute path: %w", err)
	}

	return absolutePath, nil
}

func GetAttributeFlag(usage string) *cli.StringFlag {
	return &cli.StringFlag{
		Name:    "attribute",
		Aliases: []string{"a"},
		Usage:   usage,
	}
}
