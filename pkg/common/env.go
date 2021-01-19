package common

import (
	"database/sql"
	"fmt"
)

// Env holds some application-level objects.
type Env struct {
	RDB    *sql.DB
	Config Config
}

// NewEnv creates an instance of Env.
func NewEnv(conf Config) (*Env, error) {
	rdb, err := sql.Open("mysql", conf.RDB.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the RDB: %w", err)
	}

	return &Env{
		RDB:    rdb,
		Config: conf,
	}, nil
}
