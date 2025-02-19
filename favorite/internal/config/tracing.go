package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
)

const (
	tracingServiceNameEnvName   = "TRACING_SERVICE_NAME"
	tracingCollectorHostEnvName = "TRACING_COLLECTOR_HOST"
	tracingCollectorPortEnvName = "TRACING_COLLECTOR_PORT"
)

type tracingConfig struct {
	serviceName   string
	collectorHost string
	collectorPort string
}

func NewTracingConfig() (TracingConfig, error) {
	serviceName := os.Getenv(tracingServiceNameEnvName)
	if serviceName == "" {
		return nil, fmt.Errorf("service name is not defined or empty")
	}

	collectorHost := os.Getenv(tracingCollectorHostEnvName)
	if collectorHost == "" {
		return nil, fmt.Errorf("collector host is not defined or empty")
	}

	collectorPort := os.Getenv(tracingCollectorPortEnvName)
	if collectorPort == "" {
		return nil, fmt.Errorf("collector port is not defined or empty")
	}
	_, err := strconv.Atoi(collectorPort)
	if err != nil {
		return nil, fmt.Errorf("invalid collector port: %s", err)
	}

	return &tracingConfig{
		serviceName:   serviceName,
		collectorPort: collectorPort,
		collectorHost: collectorHost,
	}, nil
}

func (c *tracingConfig) ServiceName() string {
	return c.serviceName
}

func (c *tracingConfig) CollectorAddr() string {
	return net.JoinHostPort(c.collectorHost, c.collectorPort)
}
