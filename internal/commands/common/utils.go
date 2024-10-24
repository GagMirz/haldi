package common

import (
	"fmt"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func GetFlagPath(cCtx *cli.Context) (string, error) {
	path := cCtx.String("path")
	if path == "" {
		path = "."
	}

	absolutePath, err := filepath.Abs(path)
	if err != nil {
		return "", fmt.Errorf("could not determine absolute path: %w", err)
	}

	return absolutePath, nil
}
