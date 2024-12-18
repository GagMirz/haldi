package common

import (
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

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
