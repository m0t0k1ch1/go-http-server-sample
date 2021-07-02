package db

import (
	"fmt"
	"net/url"

	ov "github.com/go-ozzo/ozzo-validation/v4"
)

const (
	parseTime = "true"
	location  = "Asia/Tokyo"
)

// Config for connecting to the DB.
type Config struct {
	Host     string `json:"host"`
	Port     uint16 `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Name     string `json:"name"`
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

// Validate validates the config.
func (conf Config) Validate() error {
	return ov.ValidateStruct(&conf,
		ov.Field(&conf.Host, ov.Required),
		ov.Field(&conf.Port, ov.Required),
		ov.Field(&conf.User, ov.Required),
		ov.Field(&conf.Name, ov.Required),
	)
}
