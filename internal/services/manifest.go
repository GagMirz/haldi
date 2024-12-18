package services

import (
	"fmt"
	"os"
	"strings"

	"haldi/internal/utils"
)

type Manifest struct {
	Name    string  `json:"name"`
	Path    string  `json:"path,omitempty"`
	Aliases []Alias `json:"aliases"`
}

var Mnfst *Manifest

func InitManifest(path string) error {
	var err error

	Mnfst, err = utils.ReadJson[Manifest](path)
	if err != nil {
		return fmt.Errorf("failed to read manifest file: %v", err)
	}

	return nil
}

func ListManifestNames() ([]string, error) {
	// List file names from manifest directory
	files, err := os.ReadDir(ManifestDirAbsPath)
	if err != nil {
		return nil, fmt.Errorf("failed to read manifest directory: %v", err)
	}

	var fileNames []string
	for _, file := range files {
		if !file.IsDir() {
			fileName := file.Name()
			if strings.HasSuffix(fileName, ".json") {
				fileName = fileName[:len(fileName)-len(".json")]
				fileNames = append(fileNames, fileName)
			}
		}
	}

	return fileNames, nil
}
