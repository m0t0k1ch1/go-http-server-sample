package common

import (
	"encoding/json"
	"fmt"
	"os"
)

const (
	defaultEnv = "dev"
)

// Config holds application settings.
type Config struct {
	Port int `json:"port"`
}

// LoadConfig loads the configuration file corresponding to APP_ENV and creates an instance of Config.
func LoadConfig() (*Config, error) {
	env := os.Getenv("APP_ENV")
	if len(env) == 0 {
		env = defaultEnv
	}

	file, err := os.Open(fmt.Sprintf("configs/%s.json", env))
	if err != nil {
		return nil, fmt.Errorf("failed to load the config file: %w", err)
	}

	var conf Config
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		return nil, fmt.Errorf("failed to decode the config: %w", err)
	}

	return &conf, nil
}
