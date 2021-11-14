package database

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type WithTxFunc func(ctx context.Context, tx *sqlx.Tx) (interface{}, error)

func WithTx(ctx context.Context, db *sqlx.DB, fn WithTxFunc) (interface{}, error) {
	t, err := db.BeginTxx(ctx, nil)
	if err != nil {
		return nil, errors.Wrap(err, "db.BeginTxx()")
	}

	res, err := fn(ctx, t)
	if err != nil {
		if errRollback := t.Rollback(); errRollback != nil {
			return nil, errors.Wrap(err, "Tx.Rollback")
		}
		return nil, errors.Wrap(err, "Tx.WithTxFunc")
	}

	if err = t.Commit(); err != nil {
		return nil, errors.Wrap(err, "Tx.Commit")
	}
	return res, nil
}
