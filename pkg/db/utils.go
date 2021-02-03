package db

import (
	"context"
	"database/sql"
)

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
