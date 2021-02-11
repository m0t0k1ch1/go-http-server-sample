package db

import (
	"context"
	"database/sql"
	"fmt"
)

// QueryParams holds parameters to build queries.
type QueryParams map[string]interface{}

// QueryPartsAndArgs returns the query parts and their arguments.
func (params QueryParams) QueryPartsAndArgs() ([]string, []interface{}) {
	parts := []string{}
	args := []interface{}{}

	for k, v := range params {
		parts = append(parts, fmt.Sprintf("%s = ?", k))
		args = append(args, v)
	}

	return parts, args
}

// Transact handles a transaction.
func Transact(ctx context.Context, db *sql.DB, txFunc func(context.Context, *sql.Tx) error) (err error) {
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}

	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			panic(p)
		} else if err != nil {
			tx.Rollback()
		} else {
			err = tx.Commit()
		}
	}()

	err = txFunc(ctx, tx)

	return
}
