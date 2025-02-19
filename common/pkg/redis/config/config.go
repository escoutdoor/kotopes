package config

import "time"

type RedisCfg interface {
	Addr() string
	DB() int
	Password() string
	MaxIdle() int
	IdleTimeout() time.Duration
	ConnTimeout() time.Duration
	TTL() time.Duration
}
