package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

var (
	grpcHostEnvName = "GRPC_HOST"
	grpcPortEnvName = "GRPC_PORT"
)

type grpcConfig struct {
	host string
	port int
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("grpc host name is empty")
	}

	portStr := os.Getenv(grpcPortEnvName)
	if portStr == "" {
		return nil, fmt.Errorf("grpc port is empty")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("conv port str into int error: %s", err)
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (c *grpcConfig) Addr() string {
	return net.JoinHostPort(c.host, strconv.Itoa(c.port))
}
