package pg

import (
	"context"

	"github.com/escoutdoor/kotopes/common/pkg/db"
	"github.com/jackc/pgx/v5/pgxpool"
)

type client struct {
	db db.DB
}

func NewClient(ctx context.Context, dsn string) (db.Client, error) {
	dbc, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, err
	}

	return &client{
		db: &pg{pool: dbc},
	}, nil
}

func (cl *client) Close() error {
	if cl.db != nil {
		cl.db.Close()
	}
	return nil
}

func (cl *client) DB() db.DB {
	return cl.db
}
