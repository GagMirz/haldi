package services

import (
	"fmt"
	"os"

	"haldi/internal/utils"
)

var ConfigPath = "/.haldi/config.json"

type Configs struct {
	Shell string `json:"shell"`
}

var Cfg *Configs

func InitConfig() error {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to find home directory")
	}

	Cfg, err = utils.ReadJson[Configs](homeDirectory + ConfigPath)
	if err != nil {
		return fmt.Errorf("failed to read haldi configuration: %v", err)
	}

	return nil
}

func OverwriteConfig() error {
	homeDirectory, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to find home directory")
	}

	err = utils.WriteJson[Configs](homeDirectory+ConfigPath, Cfg)
	if err != nil {
		return fmt.Errorf("failed to write haldi configuration: %v", err)
	}

	return nil
}
