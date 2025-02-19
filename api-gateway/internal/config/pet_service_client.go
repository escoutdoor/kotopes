package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	petClientHostEnvName = "PET_SERVICE_HOST"
	petClientPortEnvName = "PET_SERVICE_PORT"
)

type petClientConfig struct {
	host string
	port string
}

func NewPetClientConfig() (GRPCServiceClientConfig, error) {
	host := os.Getenv(petClientHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("pet service host is missing or empty")
	}

	port := os.Getenv(petClientPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("pet service port is missing or empty")
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("invalid pet service port: %w", err)
	}

	return &petClientConfig{
		host: host,
		port: port,
	}, nil
}

func (c *petClientConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}
