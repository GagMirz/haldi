package commands

import (
	"fmt"
	"os"

	"haldi/internal/services"
	"haldi/internal/utils"

	"github.com/urfave/cli/v2"
)

// TODO: Add ability to source environment from parent process
var Run = cli.Command{
	Name:        "run",
	Category:    "manifest",
	Description: "run haldi manifest alias",
	Flags:       []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		name := cCtx.Args().First()
		manifestPath := services.ManifestDir + "/" + name + ".json"

		if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
			return fmt.Errorf("manifest %s does not exist", name)
		}

		manifest, err := utils.ReadJson[services.Manifest](manifestPath)
		if err != nil {
			return fmt.Errorf("failed to read manifest: %v", err)
		}

		aliasName := cCtx.Args().Get(1)
		// Find alias by alias name in manifest
		var alias *services.Alias
		for _, a := range manifest.Aliases {
			if a.Name == aliasName {
				alias = &a
			}
		}

		if alias == nil {
			return fmt.Errorf("alias %s does not exist in manifest %s", aliasName, name)
		}

		args := cCtx.Args().Slice()[2:]
		err = alias.Run(args, manifest.Path)
		if err != nil {
			return fmt.Errorf("failed to run alias %s: %v", aliasName, err)
		}

		return nil
	},
}
