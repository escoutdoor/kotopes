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
	port string
}

func NewGRPCConfig() (GRPCConfig, error) {
	host := os.Getenv(grpcHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("grpc server host is empty")
	}

	port := os.Getenv(grpcPortEnvName)
	if len(port) == 0 {
		return nil, fmt.Errorf("grpc server port is empty")
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("invalid grpc server port: %s", err)
	}

	return &grpcConfig{
		host: host,
		port: port,
	}, nil
}

func (c *grpcConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}
