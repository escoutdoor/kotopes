package db

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type Client interface {
	DB() DB
	Close() error
}

type DB interface {
	Pinger
	SQLExecer
	Transactor
	Close()
}

type Pinger interface {
	Ping(ctx context.Context) error
}

type SQLExecer interface {
	QueryExecer
}

type Query struct {
	Name     string
	QueryRow string
}

type QueryExecer interface {
	Exec(ctx context.Context, q Query, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, q Query, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, q Query, args ...interface{}) pgx.Row
}

type Transactor interface {
	BeginTx(ctx context.Context, txOptions pgx.TxOptions) (pgx.Tx, error)
}

type TxManager interface {
	ReadCommitted(ctx context.Context, fn func(ctx context.Context) error) error
}
