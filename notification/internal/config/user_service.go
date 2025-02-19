package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

var (
	userServiceGrpcPortEnvName = "AUTH_SERVICE_GRPC_PORT"
	userServiceGrpcHostEnvName = "AUTH_SERVICE_GRPC_HOST"
)

type userServiceConfig struct {
	host string
	port string
}

func NewUserServiceConfig() (UserServiceConfig, error) {
	host := os.Getenv(userServiceGrpcHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("grpc host is not defined or empty")
	}

	port := os.Getenv(userServiceGrpcPortEnvName)
	if len(port) == 0 {
		return nil, fmt.Errorf("grpc port is not defined or empty")
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("invalid grpc port: %s", err)
	}

	return &userServiceConfig{
		host: host,
		port: port,
	}, nil
}

func (c *userServiceConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}
