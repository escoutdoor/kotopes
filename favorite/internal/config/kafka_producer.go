package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/IBM/sarama"
)

const (
	kafkaBrokersEnvName = "KAFKA_BROKERS"
)

type kafkaProducerConfig struct {
	brokers []string
}

func NewKafkaProducerConfig() (KafkaProducerConfig, error) {
	brokersStr := os.Getenv(kafkaBrokersEnvName)
	if brokersStr == "" {
		return nil, fmt.Errorf("kafka producer brokers addresses are not defined or empty")
	}
	brokers := strings.Split(brokersStr, ",")

	return &kafkaProducerConfig{
		brokers: brokers,
	}, nil
}

func (c *kafkaProducerConfig) Brokers() []string {
	return c.brokers
}

func (c *kafkaProducerConfig) Config() *sarama.Config {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	config.Producer.Return.Successes = true

	return config
}
