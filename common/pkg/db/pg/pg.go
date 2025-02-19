package pg

import (
	"context"
	"fmt"

	"github.com/escoutdoor/kotopes/common/pkg/db"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
)

const (
	TxKey string = "tx"
)

type pg struct {
	pool *pgxpool.Pool
}

func (p *pg) Close() {
	p.pool.Close()
}

func (p *pg) Exec(ctx context.Context, q db.Query, args ...interface{}) (pgconn.CommandTag, error) {
	logQuery(ctx, q)

	return p.pool.Exec(ctx, q.QueryRow, args...)
}

func (p *pg) Query(ctx context.Context, q db.Query, args ...interface{}) (pgx.Rows, error) {
	logQuery(ctx, q)

	return p.pool.Query(ctx, q.QueryRow, args...)
}

func (p *pg) QueryRow(ctx context.Context, q db.Query, args ...interface{}) pgx.Row {
	logQuery(ctx, q)

	return p.pool.QueryRow(ctx, q.QueryRow, args...)
}

func (p *pg) BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error) {
	return p.pool.BeginTx(ctx, txOptions)
}

func (p *pg) Ping(ctx context.Context) error {
	return p.pool.Ping(ctx)
}

func MakeContextTx(ctx context.Context, tx pgx.Tx) context.Context {
	return context.WithValue(ctx, TxKey, tx)
}

func logQuery(ctx context.Context, q db.Query) {
	logger.Info(
		ctx,
		fmt.Sprintf("sql: %s", q.Name),
		fmt.Sprintf("query: %s", q.QueryRow),
	)
}
