package commands

import (
	"fmt"
	"strings"

	"haldi/internal/commands/common"
	"haldi/internal/services"
	"haldi/internal/utils"

	"github.com/urfave/cli/v2"
)

var Apply = cli.Command{
	Name:        "apply",
	Category:    "config",
	Description: "applies haldi manifest",
	Flags: []cli.Flag{
		common.GetPathFlag("Path to the configuration file"),
	},
	Action: func(cCtx *cli.Context) error {
		path, err := common.GetPathFlagValue(cCtx)
		if err != nil {
			return err
		}

		// If path doesn't include haldi.json string in it, add it
		if !strings.HasSuffix(path, services.ManifestDir) {
			newSuffix := "haldi.json"
			if !strings.HasSuffix(path, "/") {
				newSuffix = "/" + newSuffix
			}

			path = path + newSuffix
		}

		// TODO: Do sanity checks on the manifest file
		manifest, err := utils.ReadJson[services.Manifest](path)
		if err != nil {
			return fmt.Errorf("failed to read manifest: %v", err)
		}

		// TODO: Add path to the manifest in Manifest struct

		err = utils.WriteJson(services.ManifestDir+"/"+manifest.Name+".json", manifest)
		if err != nil {
			return fmt.Errorf("failed to apply manifest: %v", err)
		}

		fmt.Println("Manifest applied successfully")

		return nil
	},
}
