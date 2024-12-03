package commands

import (
	"fmt"

	"haldi/internal/commands/common"
	"haldi/internal/services"

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
	Flags: []cli.Flag{
		common.GetAttributeFlag("Show attribute value"),
	},
	Action: func(cCtx *cli.Context) error {
		if cCtx.String("attribute") != "" {
			switch cCtx.String("attribute") {
			case "shell":
				fmt.Println(services.Cfg.Shell)
			default:
				return fmt.Errorf("attribute not found")
			}

			return nil
		}

		fmt.Println("Shell: ", services.Cfg.Shell)
		return nil
	},
}

var set = &cli.Command{
	Name:        "set",
	Category:    "config",
	Description: "sets haldi CLI default configurations by attribute",
	Flags: []cli.Flag{
		common.GetAttributeFlag("Set attribute value"),
	},
	Action: func(cCtx *cli.Context) error {
		if cCtx.String("attribute") != "" {
			return fmt.Errorf("please set attribute and value FE: `haldi config set shell /bin/bash`")
		}

		switch cCtx.String("attribute") {
		case "shell":
			services.Cfg.Shell = cCtx.Args().First()
		default:
			return fmt.Errorf("attribute not found")
		}

		services.OverwriteConfig()

		fmt.Println("Attribute set successfully")
		return nil
	},
}
