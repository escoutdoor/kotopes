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
	if len(dsn) == 0 {
		return nil, fmt.Errorf("postgres data source name is not defined or empty")
	}

	return &pgConfig{
		dsn: dsn,
	}, nil
}

func (c *pgConfig) DSN() string {
	return c.dsn
}
