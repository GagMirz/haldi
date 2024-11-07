package services

import (
	"fmt"
	"os"

	"haldi/internal/utils"
)

var (
	ConfigPath  = "/.haldi/config.json"
	HomeDir     string
	ManifestDir string
)

type Configs struct {
	Shell string `json:"shell"`
}

var Cfg *Configs

func InitConfig() error {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to find home directory")
	}

	HomeDir = homeDirectory
	ManifestDir = HomeDir + "/.haldi/manifests"

	Cfg, err = utils.ReadJson[Configs](HomeDir + ConfigPath)
	if err != nil {
		return fmt.Errorf("failed to read haldi default configurations: %v", err)
	}

	return nil
}

func OverwriteConfig() error {
	err := utils.WriteJson[Configs](HomeDir+ConfigPath, Cfg)
	if err != nil {
		return fmt.Errorf("failed to write haldi default configurations: %v", err)
	}

	return nil
}
