package transaction

import (
	"context"
	"fmt"

	"github.com/escoutdoor/kotopes/common/pkg/db"
	"github.com/escoutdoor/kotopes/common/pkg/db/pg"
	"github.com/jackc/pgx/v5"
)

type txManager struct {
	db db.Transactor
}

func NewTransactionManager(db db.Transactor) db.TxManager {
	return &txManager{
		db: db,
	}
}

func (m *txManager) ReadCommitted(ctx context.Context, fn func(ctx context.Context) error) error {
	txOpts := pgx.TxOptions{IsoLevel: pgx.ReadCommitted}
	return m.tr(ctx, txOpts, fn)
}

func (m *txManager) tr(ctx context.Context, opts pgx.TxOptions, fn func(ctx context.Context) error) (err error) {
	tx, ok := ctx.Value(pg.TxKey).(pgx.Tx)
	if ok {
		return fn(ctx)
	}

	tx, err = m.db.BeginTx(ctx, opts)
	if err != nil {
		return fmt.Errorf("begin tx error: %s", err)
	}

	defer func() {
		r := recover()
		if r != nil {
			err = fmt.Errorf("recover from panic: %s", err)
		}

		if err != nil {
			rollbackErr := tx.Rollback(ctx)
			if rollbackErr != nil {
				err = fmt.Errorf("rollback error: %s, rollback cause by error: %s", rollbackErr, err)
			}
			return
		}

		if err == nil {
			err = tx.Commit(ctx)
			if err != nil {
				err = fmt.Errorf("tx commit error: %s", err)
			}
		}
	}()

	err = fn(ctx)
	if err != nil {
		return fmt.Errorf("execute shenanigans inside transaction error: %s", err)
	}

	return err
}
