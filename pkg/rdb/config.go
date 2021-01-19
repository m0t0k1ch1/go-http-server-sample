package rdb

import (
	"fmt"
	"net/url"
)

const (
	parseTime = "true"
	location  = "Asia/Tokyo"
)

// Config holds the settings for connecting to the RDB.
type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
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
