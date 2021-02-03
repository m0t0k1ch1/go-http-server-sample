package common

import (
	"database/sql"
	"fmt"
)

// Env holds some application-level objects.
type Env struct {
	DB     *sql.DB
	Config Config
}

// NewEnv creates an instance of Env.
func NewEnv(conf Config) (*Env, error) {
	db, err := sql.Open("mysql", conf.DB.DSN())
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the DB: %w", err)
	}

	return &Env{
		DB:     db,
		Config: conf,
	}, nil
}
