package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"haldi/internal/commands/common"

	"github.com/urfave/cli/v2"
)

var Config = cli.Command{
	Name:        "config",
	Category:    "config",
	Description: "configs haldi CLI default behavior",
	Flags:       []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		return nil
	},
	Subcommands: []*cli.Command{
		show,
		set,
	},
}

// TODO: Move this to a separate file, as it should be available for all commands
var show = &cli.Command{
	Name:        "show",
	Category:    "config",
	Description: "shows haldi CLI default configurations",
	Flags: []cli.Flag{
		common.GetAttributeFlag("Show attribute value"),
	},
	Action: func(cCtx *cli.Context) error {
		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to find home directory")
		}

		file, err := os.Open(homeDirectory + "/.haldi/config.json")
		if err != nil {
			return fmt.Errorf("failed to open Json file: %v", err)
		}
		defer file.Close()

		var config common.DefaultConfigs
		if err := json.NewDecoder(file).Decode(&config); err != nil {
			return fmt.Errorf("failed to read Json file: %v", err)
		}

		if cCtx.String("attribute") != "" {
			switch cCtx.String("attribute") {
			case "shell":
				fmt.Println(config.Shell)
			default:
				return fmt.Errorf("attribute not found")
			}

			return nil
		}

		// Show all attributes
		fmt.Println("Shell: ", config.Shell)
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
		homeDirectory, err := os.UserHomeDir()
		if err != nil {
			return fmt.Errorf("failed to find home directory")
		}

		file, err := os.Open(homeDirectory + "/.haldi/config.json")
		if err != nil {
			return fmt.Errorf("failed to open json file: %v", err)
		}
		defer file.Close()

		var config common.DefaultConfigs
		if err := json.NewDecoder(file).Decode(&config); err != nil {
			return fmt.Errorf("failed to read json file: %v", err)
		}

		if cCtx.String("attribute") != "" {
			switch cCtx.String("attribute") {
			case "shell":
				config.Shell = cCtx.Args().First()
			default:
				return fmt.Errorf("attribute not found")
			}
		} else {
			return fmt.Errorf("please set attribute and value FE: haldi config set shell /bin/bash")
		}

		file, err = os.Create(homeDirectory + "/.haldi/config.json")
		if err != nil {
			return fmt.Errorf("failed to create json file: %v", err)
		}
		defer file.Close()

		jsonStr, err := json.MarshalIndent(config, "", " ")
		if err != nil {
			return fmt.Errorf("failed to marshal json: %v", err)
		}

		_, err = file.WriteString(string(jsonStr))
		if err != nil {
			return fmt.Errorf("failed to write to json file: %v", err)
		}

		fmt.Println("Attribute set successfully")
		return nil
	},
}
