package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

var (
	authServiceGRPCPortEnvName = "AUTH_SERVICE_GRPC_PORT"
	authServiceGRPCHostEnvName = "AUTH_SERVICE_GRPC_HOST"
)

type authServiceGRPCConfig struct {
	host string
	port int
}

func NewAuthServiceGRPCConfig() (AuthServiceGRPCConfig, error) {
	host := os.Getenv(authServiceGRPCHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("auth service grpc host is empty")
	}

	portStr := os.Getenv(authServiceGRPCPortEnvName)
	if len(portStr) == 0 {
		return nil, fmt.Errorf("auth service grpc port is empty")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("auth service grpc port should be an integer")
	}

	return &authServiceGRPCConfig{
		host: host,
		port: port,
	}, nil
}

func (c *authServiceGRPCConfig) Addr() string {
	return net.JoinHostPort(c.host, strconv.Itoa(c.port))
}
