package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
)


const configFileName = ".gatorconfig.json"


type Config struct {
	DbUrl string `json:"db_url,omitempty"`
	CurrentUserName string `json:"current_user_name,omitempty"`
}


func (c *Config) SetUser(name string) error {
	c.CurrentUserName = name
	//marshals but keeps indents so easier to read the json
	jsonBytes, err := json.MarshalIndent(c, "", "\t")
	if err != nil {
		return fmt.Errorf("json marshalling: %w", err)
	}

	filePath, err := getFilePath()
	if err != nil {
		return err
	}
	//writes json to file.
	if err = os.WriteFile(filePath, jsonBytes, 0644); err != nil{
		return fmt.Errorf("writing file: %w", err)
	}
	return nil
}


func getFilePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("getting file path: %w", err)
	}
	return filepath.Join(homeDir, configFileName), nil
}


func Read() (Config, error) {
	filePath, err := getFilePath()
	if err != nil {
		return Config{}, err
	}

	jsonBytes, err := os.ReadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("reading file: %w", err)
	}

	var config Config
	err = json.Unmarshal(jsonBytes, &config)
	if err != nil {
		return Config{}, fmt.Errorf("json unmarshalling: %w", err)
	}
	return config, nil
}
