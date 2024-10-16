package manager

import (
	"encoding/json"
	"fmt"
	"os"

	confmap "haldi/internal/config-map"
)

type Manager struct {
	ConfigPath string            `json:"configPath"`
	ConfigMap  *confmap.ConfigMap `json:"configMap"`
}

func (mngr *Manager) ReadConfig() error {
	file, err := os.Open(mngr.ConfigPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&mngr.ConfigMap)
	if err != nil {
		return err
	}

	fmt.Printf("%v", mngr.ConfigMap)

	return nil
}

var mngr *Manager

func InitManager(configPath string) error {
	mngr = &Manager{
		ConfigPath: configPath,
		ConfigMap:  &confmap.ConfigMap{},
	}

	return mngr.ReadConfig()
}
