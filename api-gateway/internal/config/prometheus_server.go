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
	port int
}

func NewPrometheusConfig() (PrometheusServerConfig, error) {
	host := os.Getenv(prometheusHostEnvName)
	if len(host) == 0 {
		return nil, fmt.Errorf("prometheus host is empty")
	}

	portStr := os.Getenv(prometheusPortEnvName)
	if len(portStr) == 0 {
		return nil, fmt.Errorf("prometheus port is empty")
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, fmt.Errorf("prometheus port should be an integer")
	}

	return &prometheusConfig{
		host: host,
		port: port,
	}, nil
}

func (c *prometheusConfig) Addr() string {
	return net.JoinHostPort(c.host, strconv.Itoa(c.port))
}
