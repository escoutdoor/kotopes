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
			return fmt.Errorf("config file doesn't exist")
		}
		return fmt.Errorf("stat config file error: %s", err)
	}

	err = godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("load config error: %s", err)
	}
	return nil
}

type GRPCConfig interface {
	Addr() string
}

type PGConfig interface {
	DSN() string
}

type TokenConfig interface {
	SecretKey() []byte
	TTL() time.Duration
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
