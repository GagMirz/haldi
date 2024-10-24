package commands

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"haldi/internal/commands/common"

	"github.com/urfave/cli/v2"
)

var Init = cli.Command{
	Name:        "init",
	Category:    "config",
	Description: "creates .haldi.json config file with basic structure",
	Flags: []cli.Flag{
		common.GetNameFlag("Configuration name, will be used for calling commands"),
		common.GetPathFlag("Path, where to create .haldi.json, defaults to '.'"),
	},
	Action: func(cCtx *cli.Context) error {
		fileName := ".haldi.json"

		name := cCtx.String("name")

		if name == "" {
			return fmt.Errorf("name cannot be empty, please provide a name using the -n flag")
		}

		absolutePath, err := common.GetFlagPath(cCtx)
		if err != nil {
			return err
		}

		if err := os.MkdirAll(absolutePath, os.ModePerm); err != nil {
			return fmt.Errorf("could not create directory: %w", err)
		}

		filePath := filepath.Join(absolutePath, fileName)

		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			return fmt.Errorf("config file already exists")
		}

		file, err := os.Create(filePath)
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

		fmt.Printf("Created %s template file in %s\n", fileName, absolutePath)
		return nil
	},
}
