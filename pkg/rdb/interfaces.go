package rdb

import (
	"context"
	"database/sql"
)

// Executer is an interface to execute queries.
type Executer interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

// Scanner is an interface to scan a single row.
type Scanner interface {
	Scan(dest ...interface{}) error
}
