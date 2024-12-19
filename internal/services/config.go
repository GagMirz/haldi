package services

import (
	"fmt"
	"os"

	"haldi/internal/utils"
)

var (
	ConfigPath         = "/.haldi/config.json"
	ManifestPath       = "/.haldi/manifests"
	HomeDir            string
	ManifestDirAbsPath string
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
	ManifestDirAbsPath = HomeDir + ManifestPath

	if _, err := os.Stat(ManifestDirAbsPath); os.IsNotExist(err) {
		err := os.MkdirAll(ManifestDirAbsPath, os.ModePerm)
		if err != nil {
			return fmt.Errorf("failed to create manifest dir: %v", err)
		}
	}

	// Check if config files exists
	configFileAbsPath := HomeDir + ConfigPath
	if _, err := os.Stat(configFileAbsPath); os.IsNotExist(err) {
		// Create default config file
		err := utils.WriteJson[Configs](configFileAbsPath, &Configs{Shell: "bash"})
		if err != nil {
			return fmt.Errorf("failed to create haldi default configurations: %v", err)
		}
	}

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
