package common

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/rdb"
)

const (
	defaultEnv = "dev"
)

// Config holds the settings for the main application.
type Config struct {
	Port int        `json:"port"`
	RDB  rdb.Config `json:"rdb"`
}

// LoadConfig loads the configuration file corresponding to APP_ENV and creates an instance of Config.
func LoadConfig() (Config, error) {
	env := os.Getenv("APP_ENV")
	if len(env) == 0 {
		env = defaultEnv
	}

	file, err := os.Open(fmt.Sprintf("configs/%s.json", env))
	if err != nil {
		return Config{}, fmt.Errorf("failed to load the config file: %w", err)
	}

	var conf Config
	if err := json.NewDecoder(file).Decode(&conf); err != nil {
		return Config{}, fmt.Errorf("failed to decode the config: %w", err)
	}

	return conf, nil
}
