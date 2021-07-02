package app

import (
	"fmt"

	ov "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/kayac/go-config"

	"github.com/m0t0k1ch1/go-http-server-sample/pkg/db"
)

// Config for the main application.
type Config struct {
	Port uint16    `json:"port"`
	DB   db.Config `json:"db"`
}

// LoadConfig loads the config.
func LoadConfig(path string) (Config, error) {
	var conf Config
	if err := config.LoadWithEnvJSON(&conf, path); err != nil {
		return Config{}, fmt.Errorf("failed to load the config: %w", err)
	}

	return conf, nil
}

// Validate validates the config.
func (conf Config) Validate() error {
	return ov.ValidateStruct(&conf,
		ov.Field(&conf.Port, ov.Required),
		ov.Field(&conf.DB, ov.Required),
	)
}
