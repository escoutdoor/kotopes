package config

import (
	"fmt"
	"os"

	"github.com/IBM/sarama"
	"github.com/joho/godotenv"
)

func Load(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return fmt.Errorf("file doesn't exist: %s", err)
		}
		return err
	}

	err = godotenv.Load(path)
	if err != nil {
		return fmt.Errorf("load config: %s", err)
	}
	return nil
}

type KafkaConsumerConfig interface {
	Brokers() []string
	GroupID() string
	Config() *sarama.Config
}

type FavoriteServiceConfig interface {
	Topic() string
}

type MailConfig interface {
	Host() string
	Port() int
	From() string
	Password() string
}

type UserServiceConfig interface {
	Addr() string
}
