package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"haldi/internal/commands/common"
	"haldi/internal/services"
	"haldi/internal/utils"

	"github.com/urfave/cli/v2"
)

var ManifestFileName = "haldi.json"

var Init = cli.Command{
	Name:        "init",
	Category:    "config",
	Description: "creates .haldi.json config file with basic structure",
	Flags: []cli.Flag{
		common.GetNameFlag("Configuration name, will be used for calling commands"),
		common.GetPathFlag("Path, where to create .haldi.json, defaults to '.'"),
	},
	Action: func(cCtx *cli.Context) error {
		name, err := common.GetNameFlagValue(cCtx)
		if err != nil {
			return err
		}

		// TODO: make name not required, use repository name as default
		if name == "" {
			return fmt.Errorf("name is required, please provide name with -n flag")
		}

		absolutePath, err := common.GetPathFlagValue(cCtx)
		if err != nil {
			return err
		}

		if err := os.MkdirAll(absolutePath, os.ModePerm); err != nil {
			return fmt.Errorf("could not create directory: %w", err)
		}

		filePath := filepath.Join(absolutePath, ManifestFileName)

		if _, err := os.Stat(filePath); !os.IsNotExist(err) {
			var overwrite string
			fmt.Printf("File %s already exists, do you want to overwrite it? (yes/no): ", filePath)
			fmt.Scanf("%s", &overwrite)

			if overwrite != "yes" {
				return fmt.Errorf("command aborted")
			}

			if err := os.Remove(filePath); err != nil {
				return fmt.Errorf("could not remove existing file: %w", err)
			}
		}

		manifest := services.Manifest{
			Name: name,
		}

		err = utils.WriteJson(filePath, &manifest)
		if err != nil {
			return fmt.Errorf("failed to create manifest file: %w", err)
		}

		fmt.Printf("Created %s template file in %s\n", ManifestFileName, absolutePath)
		return nil
	},
}
