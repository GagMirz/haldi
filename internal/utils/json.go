package utils

import (
	"encoding/json"
	"fmt"
	"os"
)

func ReadJson[T any](path string) (*T, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open json file: %v", err)
	}
	defer file.Close()

	var jsonStruct T
	if err := json.NewDecoder(file).Decode(&jsonStruct); err != nil {
		return nil, fmt.Errorf("failed to read json file: %v", err)
	}

	return &jsonStruct, nil
}

func WriteJson[T any](path string, jsonStruct *T) error {
	file, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("failed to create json file: %v", err)
	}
	defer file.Close()

	jsonStr, err := json.MarshalIndent(jsonStruct, "", " ")
	if err != nil {
		return fmt.Errorf("failed to marshal json: %v", err)
	}

	_, err = file.WriteString(string(jsonStr))
	if err != nil {
		return fmt.Errorf("failed to write to json file: %v", err)
	}

	return nil
}
