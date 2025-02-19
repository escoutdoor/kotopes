package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

var (
	redisHostEnvName        = "REDIS_HOST"
	redisPortEnvName        = "REDIS_PORT"
	redisPasswordEnvName    = "REDIS_PASSWORD"
	redisDBEnvName          = "REDIS_DB"
	redisMaxIdleEnvName     = "REDIS_MAX_IDLE"
	redisIdleTimeoutEnvName = "REDIS_IDLE_TIMEOUT"
	redisConnTimeoutEnvName = "REDIS_CONNECTION_TIMEOUT"
	redisTTLEnvName         = "REDIS_TTL"
)

type redisConfig struct {
	host string
	port int

	password string
	db       int

	maxIdle     int
	idleTimeout time.Duration
	connTimeout time.Duration

	ttl time.Duration
}

func NewRedisConfig() (RedisConfig, error) {
	host := os.Getenv(redisHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("redis host is empty")
	}

	portStr := os.Getenv(redisPortEnvName)
	if portStr == "" {
		return nil, fmt.Errorf("redis port is empty")
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("convert port string to int error: %s", err)
	}

	password := os.Getenv(redisPasswordEnvName)

	dbStr := os.Getenv(redisDBEnvName)
	db, err := strconv.Atoi(dbStr)
	if err != nil {
		return nil, fmt.Errorf("convert db string to int error: %s", err)
	}

	maxIdleStr := os.Getenv(redisMaxIdleEnvName)
	if maxIdleStr == "" {
		return nil, fmt.Errorf("redis max idle is empty")
	}
	maxIdle, err := strconv.Atoi(maxIdleStr)
	if err != nil {
		return nil, fmt.Errorf("convert max idle string to int error: %s", err)
	}

	idleTimeoutStr := os.Getenv(redisIdleTimeoutEnvName)
	if idleTimeoutStr == "" {
		return nil, fmt.Errorf("idle timeout is empty")
	}
	idleTimeout, err := time.ParseDuration(idleTimeoutStr)
	if err != nil {
		return nil, fmt.Errorf("parse idle timeout error: %s", err)
	}

	connTimeoutStr := os.Getenv(redisConnTimeoutEnvName)
	if connTimeoutStr == "" {
		return nil, fmt.Errorf("connection timeout is empty")
	}
	connTimeout, err := time.ParseDuration(connTimeoutStr)
	if err != nil {
		return nil, fmt.Errorf("parse connection timeout error: %s", err)
	}

	ttlStr := os.Getenv(redisTTLEnvName)
	if ttlStr == "" {
		return nil, fmt.Errorf("redis cache ttl is empty")
	}
	ttl, err := time.ParseDuration(ttlStr)
	if err != nil {
		return nil, fmt.Errorf("parse redis time to live error: %s", err)
	}

	return &redisConfig{
		host:        host,
		port:        port,
		password:    password,
		db:          db,
		maxIdle:     maxIdle,
		idleTimeout: idleTimeout,
		connTimeout: connTimeout,
		ttl:         ttl,
	}, nil
}

func (c *redisConfig) Addr() string {
	return net.JoinHostPort(c.host, strconv.Itoa(c.port))
}

func (c *redisConfig) DB() int {
	return c.db
}

func (c *redisConfig) Password() string {
	return c.password
}

func (c *redisConfig) MaxIdle() int {
	return c.maxIdle
}

func (c *redisConfig) IdleTimeout() time.Duration {
	return c.idleTimeout
}

func (c *redisConfig) ConnTimeout() time.Duration {
	return c.connTimeout
}

func (c *redisConfig) TTL() time.Duration {
	return c.ttl
}
