package commands

import (
	"fmt"
	"os"

	"haldi/internal/services"

	"github.com/urfave/cli/v2"
)

var Remove = cli.Command{
	Name:        "remove",
	Category:    "manifest",
	Description: "remove haldi manifest",
	Flags:       []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		name := cCtx.Args().First()
		manifestPath := services.ManifestDir + "/" + name + ".json"

		if _, err := os.Stat(manifestPath); os.IsNotExist(err) {
			return fmt.Errorf("manifest %s does not exist", name)
		}

		err := os.Remove(manifestPath)
		if err != nil {
			return fmt.Errorf("failed to remove manifest: %v", err)
		}

		return nil
	},
}
