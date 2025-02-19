package config

import (
	"os"

	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/joho/godotenv"
)

func Load(path string) error {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errwrap.Wrapf("config file (%s) does not exist", err, path)
	}

	err := godotenv.Load(path)
	if err != nil {
		return errwrap.Wrap("load config error", err)
	}

	return nil
}

type HTTPServerConfig interface {
	Addr() string
}

type GRPCServiceClientConfig interface {
	Addr() string
}

type PrometheusServerConfig interface {
	Addr() string
}

type TracingConfig interface {
	ServiceName() string
	CollectorAddr() string
}
