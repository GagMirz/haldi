package commands

import (
	"encoding/json"
	"fmt"
	"os"

	"haldi/internal/commands/common"

	"github.com/urfave/cli/v2"
)

var Init = cli.Command{
	Name:        "init",
	Category:    "config",
	Description: "creates .haldi.json config file with basic structure",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:    "name",
			Aliases: []string{"n"},
			Usage:   "Configuration name",
		},
	},
	Action: func(cCtx *cli.Context) error {
		fileName := ".haldi.json"

		name := cCtx.String("name")

		if name == "" {
			return fmt.Errorf("name cannot be empty, please provide a name using the -n flag")
		}

		if _, err := os.Stat(fileName); !os.IsNotExist(err) {
			return fmt.Errorf("config file already exists")
		}

		file, err := os.Create(fileName)
		if err != nil {
			return fmt.Errorf("failed to create file: %w", err)
		}
		defer file.Close()

		jsonStr, err := json.MarshalIndent(common.InitJsonContent{
			Name: name,
		}, "", " ")
		if err != nil {
			return fmt.Errorf("failed to marshal json: %w", err)
		}

		_, err = file.WriteString(string(jsonStr))
		if err != nil {
			return fmt.Errorf("failed to write to file: %w", err)
		}

		fmt.Printf("Created .haldi.json template file: %s\n", fileName)
		return nil
	},
}
