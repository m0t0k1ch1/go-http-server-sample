package app

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config holds application settings.
type Config struct {
	Port int `json:"port"`
}

// LoadConfig loads the configuration file and creates an instance of Config
func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	var conf Config
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		return nil, fmt.Errorf("failed to load config: %w", err)
	}

	return &conf, nil
}
