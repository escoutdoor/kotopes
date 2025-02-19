package config

import (
	"fmt"
	"os"
)

const (
	favoriteKafkaTopicEnvName = "KAFKA_FAVORITE_TOPIC"
)

type notificationServiceConfig struct {
	topic string
}

func NewNotificationServiceConfig() (NotificationServiceConfig, error) {
	topic := os.Getenv(favoriteKafkaTopicEnvName)
	if topic == "" {
		return nil, fmt.Errorf("favorite kafka topic is not defined or empty")
	}

	return &notificationServiceConfig{
		topic: topic,
	}, nil
}

func (c *notificationServiceConfig) Topic() string {
	return c.topic
}
