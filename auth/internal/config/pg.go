package config

import (
	"fmt"
	"os"
)

var (
	pgDSNEnvName = "PG_DSN"
)

type pgConfig struct {
	dsn string
}

func NewPGConfig() (PGConfig, error) {
	dsn := os.Getenv(pgDSNEnvName)
	if dsn == "" {
		return nil, fmt.Errorf("pg dsn in empty")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (c *pgConfig) DSN() string {
	return c.dsn
}
