package db

import (
	"fmt"
	"net/url"

	"github.com/m0t0k1ch1/go-envparser"
)

const (
	parseTime = "true"
	location  = "Asia/Tokyo"
)

// Config for connecting to the DB.
type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

// LoadConfigFromEnv loads the config from environment variables.
func LoadConfigFromEnv() (Config, error) {
	var err error
	parse := func(k string, v interface{}) {
		if err != nil {
			return
		}
		err = envparser.Parse(k, v)
	}

	var conf Config
	parse("APP_DB_HOST", &conf.Host)
	parse("APP_DB_PORT", &conf.Port)
	parse("APP_DB_USER", &conf.User)
	parse("APP_DB_PASSWORD", &conf.Password)
	parse("APP_DB_NAME", &conf.Name)
	if err != nil {
		return Config{}, fmt.Errorf("failed to load the config from environment variables: %w", err)
	}

	return conf, nil
}

// DSN returns the data source name for database/sql.
func (conf Config) DSN() string {
	q := url.Values{}
	q.Add("parseTime", parseTime)
	q.Add("loc", location)

	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?%s",
		conf.User, conf.Password, conf.Host, conf.Port, conf.Name, q.Encode(),
	)
}
