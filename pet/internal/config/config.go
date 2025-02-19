package config

import (
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func Load(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file doesn't exist: %s", err)
		}
		return err
	}

	err = godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("load config: %s", err)
	}
	return nil
}

type AuthServiceGRPCConfig interface {
	Addr() string
}

type GRPCConfig interface {
	Addr() string
}

type PGConfig interface {
	DSN() string
}

type RedisConfig interface {
	Addr() string
	DB() int
	Password() string
	MaxIdle() int
	IdleTimeout() time.Duration
	ConnTimeout() time.Duration
	TTL() time.Duration
}

type PrometheusConfig interface {
	Addr() string
}

type TracingConfig interface {
	ServiceName() string
	CollectorAddr() string
}
