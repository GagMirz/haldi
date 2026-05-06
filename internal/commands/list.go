package commands

import (
	"fmt"

	"anvil/internal/services"

	"github.com/urfave/cli/v2"
)

var List = cli.Command{
	Name:        "list",
	Category:    "manifest",
	Description: "lists anvil manifests",
	Flags:       []cli.Flag{},
	Action: func(cCtx *cli.Context) error {
		manifestNames, err := services.ListManifestNames()
		if err != nil {
			return fmt.Errorf("failed to list manifests: %v", err)
		}

		fmt.Println("Amount of manifests:", len(manifestNames))
		fmt.Println("List of manifests:")

		for i, manifestName := range manifestNames {
			fmt.Printf("%d) %s\n", i+1, manifestName)
		}

		return nil
	},
}
