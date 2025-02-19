package config

import (
	"fmt"
	"os"
)

const (
	favoriteKafkaTopicEnvName = "KAFKA_FAVORITE_TOPIC"
)

type favoriteServiceConfig struct {
	topic string
}

func NewFavoriteServiceConfig() (FavoriteServiceConfig, error) {
	topic := os.Getenv(favoriteKafkaTopicEnvName)
	if topic == "" {
		return nil, fmt.Errorf("kafka topic is not defined or empty")
	}

	return &favoriteServiceConfig{
		topic: topic,
	}, nil
}

func (c *favoriteServiceConfig) Topic() string {
	return c.topic
}
