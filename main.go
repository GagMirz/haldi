package main

import (
	"fmt"
	"os"

	"haldi/internal/commands"
	"haldi/internal/services"

	"github.com/urfave/cli/v2"
)

// TODO: Unify error messages
// TODO: Add license to repository
// TODO: Add command to check manifest directory for corrupted manifests and non-json files
// TODO: Add feature to restore manifest file from applied manifests
// TODO: Add global manifest option, so manifest won't be tied to project/repository
// TODO: Add errorHint attribute to manifest, so user can get hints on how to fix the error

func main() {
	services.InitConfig()

	app := &cli.App{
		Name:      "haldi",
		Usage:     "Repository accessibility extender",
		Version:   "v0.0.1",
		UsageText: "haldi command [command options]",
		Commands: []*cli.Command{
			&commands.Init,
			&commands.Config,
			&commands.Apply,
			&commands.List,
			&commands.Remove,
			&commands.Run,
		},
		CommandNotFound: func(cCtx *cli.Context, command string) {
			fmt.Fprintf(cCtx.App.Writer, "Command %q not found.\n", command)
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}
}
