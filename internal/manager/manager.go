package manager

import (
	"encoding/json"
	"fmt"
	"os"

	confmap "haldi/internal/config-map"
)

func ReadConfig(configPath string) error {
	file, err := os.Open(configPath)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	configMap := &confmap.ConfigMap{}

	err = decoder.Decode(configMap)
	if err != nil {
		return err
	}

	fmt.Printf("%v", configMap)

	return nil
}
