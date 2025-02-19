package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	favoriteClientHostEnvName = "FAVORITE_SERVICE_HOST"
	favoriteClientPortEnvName = "FAVORITE_SERVICE_PORT"
)

type favoriteClientConfig struct {
	host string
	port string
}

func NewFavoriteClientConfig() (GRPCServiceClientConfig, error) {
	host := os.Getenv(favoriteClientHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("favorite service host is missing or empty")
	}

	port := os.Getenv(favoriteClientPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("favorite service port is missing or empty")
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("invalid favorite service port: %w", err)
	}

	return &favoriteClientConfig{
		host: host,
		port: port,
	}, nil
}

func (c *favoriteClientConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}
