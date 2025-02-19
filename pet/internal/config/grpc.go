package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

var (
	grpcPortEnvName = "GRPC_PORT"
	grpcHostEnvName = "GRPC_HOST"
)

type grpcConfig struct {
	host string
	port int
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("grpc host is empty")
	}

	portStr := os.Getenv(grpcPortEnvName)
	if len(portStr) == 0 {
		return nil, fmt.Errorf("grpc port is empty")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("grpc port should be an integer")
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (c *grpcConfig) Addr() string {
	return net.JoinHostPort(c.host, strconv.Itoa(c.port))
}
