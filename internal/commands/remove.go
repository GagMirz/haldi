package commands

import (
	"fmt"
	"os"
	"path/filepath"

	"anvil/internal/services"
	"anvil/internal/utils"

	"github.com/urfave/cli/v2"
)

var Remove = cli.Command{
	Name:        "remove",
	Category:    "manifest",
	Description: "remove anvil manifest",
	Flags:       []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		name := cCtx.Args().Get(0)

		if name == "" {
			path, err := filepath.Abs(".")
			if err != nil {
				return err
			}

			path += "/anvil.json"

			manifest, err := utils.ReadJson[services.Manifest](path)
			if err != nil {
				return fmt.Errorf("failed to read manifest: %v", err)
			}

			name = manifest.Name
		}

		manifestPath := services.ManifestDirAbsPath + "/" + name + ".json"

		if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
			return fmt.Errorf("manifest %s does not exist", name)
		}

		err := os.Remove(manifestPath)
		if err != nil {
			return fmt.Errorf("failed to remove manifest: %v", err)
		}

		fmt.Printf("Manifest %s removed\n", name)

		return nil
	},
}
