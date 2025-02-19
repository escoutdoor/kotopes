package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	authClientHostEnvName = "AUTH_SERVICE_HOST"
	authClientPortEnvName = "AUTH_SERVICE_PORT"
)

type authClientConfig struct {
	host string
	port string
}

func NewAuthClientConfig() (GRPCServiceClientConfig, error) {
	host := os.Getenv(authClientHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("auth service host is missing or empty")
	}

	port := os.Getenv(authClientPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("auth service port is missing or empty")
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("invalid auth service port: %w", err)
	}

	return &authClientConfig{
		host: host,
		port: port,
	}, nil
}

func (c *authClientConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}
