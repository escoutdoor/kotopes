package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

var (
	prometheusPortEnvName = "PROMETHEUS_SERVER_PORT"
	prometheusHostEnvName = "PROMETHEUS_SERVER_HOST"
)

type prometheusConfig struct {
	host string
	port string
}

func NewPrometheusConfig() (PrometheusConfig, error) {
	host := os.Getenv(prometheusHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("prometheus host is not defined")
	}

	port := os.Getenv(prometheusPortEnvName)
	if len(port) == 0 {
		return nil, fmt.Errorf("prometheus port is not defined")
	}
	_, err := strconv.Atoi(port)
	if err != nil {
		return nil, fmt.Errorf("invalid prometheus port: %s", err)
	}

	return &prometheusConfig{
		host: host,
		port: port,
	}, nil
}

func (c *prometheusConfig) Addr() string {
	return net.JoinHostPort(c.host, c.port)
}
