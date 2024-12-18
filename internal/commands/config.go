package commands

import (
	"fmt"

	"haldi/internal/services"
	"haldi/internal/utils"

	"github.com/urfave/cli/v2"
)

var Config = cli.Command{
	Name:        "config",
	Category:    "config",
	Description: "configs haldi CLI default behavior",
	Flags:       []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		// TODO: Add help message for subcommands
		return nil
	},
	Subcommands: []*cli.Command{
		show,
		set,
	},
}

var show = &cli.Command{
	Name:        "show",
	Category:    "config",
	Description: "shows haldi CLI default configurations",
	Flags:       []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		attribute := cCtx.Args().Get(0)
		switch attribute {
		case "shell":
			fmt.Println(services.Cfg.Shell)
		case "":
			// List all available attributes
			fmt.Println("Shell: ", services.Cfg.Shell)
		default:
			return fmt.Errorf("attribute not found")
		}

		return nil
	},
}

var set = &cli.Command{
	Name:        "set",
	Category:    "config",
	Description: "sets haldi CLI default configurations by attribute",
	Flags:       []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		attribute := cCtx.Args().Get(0)

		if attribute == "" {
			return fmt.Errorf("please set attribute and value FE: `haldi config set shell /bin/bash`")
		}

		switch attribute {
		case "shell":
			value := cCtx.Args().Get(1)

			if utils.Contains([]string{"zsh", "bash", "sh", "ksh", "dash"}, value) {
				return fmt.Errorf("unsupported shell")
			}

			services.Cfg.Shell = cCtx.Args().Get(1)
		default:
			return fmt.Errorf("attribute not found")
		}

		services.OverwriteConfig()

		fmt.Println("Attribute set successfully")
		return nil
	},
}
