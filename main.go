package main

import (
	"fmt"
	"os"

	"haldi/internal/commands"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:      "haldi",
		Usage:     "Repository accessibility extender",
		Version:   "v0.0.1",
		UsageText: "haldi command [command options]",
		Commands: []*cli.Command{
			&commands.Init,
			&commands.Config,
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
