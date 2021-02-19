package common

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/m0t0k1ch1/go-envparser"
	"github.com/m0t0k1ch1/go-http-server-sample/pkg/db"
)

const (
	defaultEnv = "dev"
)

// Config for the main application.
type Config struct {
	Port int       `json:"port"`
	DB   db.Config `json:"db"`
}

// LoadConfig loads the config.
func LoadConfig(path string) (Config, error) {
	if path == "" {
		return LoadConfigFromEnv()
	}

	return LoadConfigFromFile(path)
}

// LoadConfigFromFile loads the config from a file.
func LoadConfigFromFile(path string) (Config, error) {
	file, err := os.Open(path)
	if err != nil {
		return Config{}, fmt.Errorf("failed to open the config file: %w", err)
	}

	var conf Config
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		return Config{}, fmt.Errorf("failed to decode the config: %w", err)
	}

	return conf, nil
}

// LoadConfigFromEnv loads the config from environment variables.
func LoadConfigFromEnv() (Config, error) {
	var conf Config
	if err := envparser.Parse("APP_PORT", &conf.Port); err != nil {
		return Config{}, fmt.Errorf("failed to load the config from environment variables: %w", err)
	}

	dbConf, err := db.LoadConfigFromEnv()
	if err != nil {
		return Config{}, err
	}

	conf.DB = dbConf

	return conf, nil
}
