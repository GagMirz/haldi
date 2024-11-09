package main

import (
	"fmt"
	"os"

	"haldi/internal/commands"
	"haldi/internal/services"

	"github.com/urfave/cli/v2"
)

// TODO: Unify error messages
// TODO: Add command to check manifest directory for corrupted manifests and non-json files

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
