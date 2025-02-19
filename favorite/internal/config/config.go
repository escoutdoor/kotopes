package config

import (
	"os"

	"github.com/IBM/sarama"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/joho/godotenv"
)

func Load(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return errwrap.Wrap("file does not exist", err)
		}
		return err
	}

	err = godotenv.Load(path)
	if err != nil {
		return errwrap.Wrap("load", err)
	}
	return nil
}

type PGConfig interface {
	DSN() string
}

type GRPCConfig interface {
	Addr() string
}

type KafkaProducerConfig interface {
	Brokers() []string
	Config() *sarama.Config
}

type NotificationServiceConfig interface {
	Topic() string
}

type PetServiceConfig interface {
	Addr() string
}

type PrometheusConfig interface {
	Addr() string
}

type TracingConfig interface {
	ServiceName() string
	CollectorAddr() string
}
