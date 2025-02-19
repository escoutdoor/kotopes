package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	httpServerHostEnvName = "HTTP_SERVER_HOST"
	httpServerPortEnvName = "HTTP_SERVER_PORT"
)

type httpServerConfig struct {
	host string
	port string
}

var _ (HTTPServerConfig) = (*httpServerConfig)(nil)

func NewHTTPServerConfig() (HTTPServerConfig, error) {
	host := os.Getenv(httpServerHostEnvName)
	if host == "" {
		return nil, fmt.Errorf("HTTP server host is missing or empty")
	}

	port := os.Getenv(httpServerPortEnvName)
	if port == "" {
		return nil, fmt.Errorf("HTTP server port is missing or empty")
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("invalid HTTP server port: %w", err)
	}

	return &httpServerConfig{
		host: host,
		port: port,
	}, nil
}

func (c *httpServerConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}
