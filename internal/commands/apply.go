package commands

import (
	"fmt"
	"os"
	"strings"

	"haldi/internal/services"
	"haldi/internal/utils"

	"github.com/urfave/cli/v2"
)

var Apply = cli.Command{
	Name:        "apply",
	Category:    "config",
	Description: "applies haldi manifest",
	Flags:       []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		path := cCtx.Args().Get(0)
		if path == "" {
			path = "."
		}

		if !strings.HasSuffix(path, services.ManifestDir) {
			newSuffix := "haldi.json"
			if !strings.HasSuffix(path, "/") {
				newSuffix = "/" + newSuffix
			}

			path = path + newSuffix
		}

		// TODO: Do sanity checks on the manifest file
		// 1) Check aliases for duplicates
		// 2) Check if all aliases are valid
		manifest, err := utils.ReadJson[services.Manifest](path)
		if err != nil {
			return fmt.Errorf("failed to read manifest: %v", err)
		}

		projectPath := strings.TrimSuffix(path, "/haldi.json")
		manifest.Path = projectPath

		appliedManifestPath := services.ManifestDir + "/" + manifest.Name + ".json"

		if _, err := os.Stat(appliedManifestPath); !os.IsNotExist(err) {
			confirmation := utils.RequestConfirmation(fmt.Sprintf("File %s already exists, do you want to overwrite it? (yes/no): ", appliedManifestPath))
			if !confirmation {
				return fmt.Errorf("command aborted")
			}

			if err := os.Remove(appliedManifestPath); err != nil {
				return fmt.Errorf("could not remove existing file: %w", err)
			}
		}

		err = utils.WriteJson(appliedManifestPath, manifest)
		if err != nil {
			return fmt.Errorf("failed to apply manifest: %v", err)
		}

		fmt.Println("Manifest applied successfully")

		return nil
	},
}
