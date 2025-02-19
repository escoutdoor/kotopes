package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

var (
	petServicePortEnvName = "PET_SERVICE_GRPC_PORT"
	petServiceHostEnvName = "PET_SERVICE_GRPC_HOST"
)

type petServiceConfig struct {
	host string
	port string
}

func NewPetServiceConfig() (PetServiceConfig, error) {
	host := os.Getenv(petServiceHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("pet service grpc host is not defined or empty")
	}

	port := os.Getenv(petServicePortEnvName)
	if len(port) == 0 {
		return nil, fmt.Errorf("pet service grpc port is not defined or empty")
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("invalid pet service grpc port: %s", err)
	}

	return &petServiceConfig{
		host: host,
		port: port,
	}, nil
}

func (c *petServiceConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}
