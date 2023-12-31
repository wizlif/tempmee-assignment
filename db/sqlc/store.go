package db

import (
	"context"
	"database/sql"
	"fmt"
)

// store provides all functions to execute db queries and transactions
type Store interface {
	Querier
}

// store provides all functions to execute db queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

// NewStore creates new store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}

// execTx executes a function within a database transaction
func (store *SQLStore) execTxn(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.BeginTx(ctx, nil)

	if err != nil {
		return err
	}

	q := New(tx)

	err = fn(q)

	if err != nil {
		if rbErr := tx.Rollback(); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", err, rbErr)
		}

		return err
	}

	return tx.Commit()
}
