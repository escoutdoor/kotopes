package config

import (
	"fmt"
	"os"
	"strings"

	"github.com/IBM/sarama"
)

const (
	kafkaFavoriteBrokersEnvName = "KAFKA_FAVORITE_BROKERS"
	kafkaFavoriteGroupIDEnvName = "KAFKA_FAVORITE_GROUP_ID"
)

type kafkaFavoriteConsumerConfig struct {
	brokers []string
	groupID string
}

func NewKafkaFavoriteConsumerConfig() (KafkaConsumerConfig, error) {
	brokersStr := os.Getenv(kafkaFavoriteBrokersEnvName)
	if brokersStr == "" {
		return nil, fmt.Errorf("kafka brokers addresses are not defined or empty")
	}

	brokers := strings.Split(brokersStr, ",")

	groupID := os.Getenv(kafkaFavoriteGroupIDEnvName)
	if groupID == "" {
		return nil, fmt.Errorf("kafka consumer group id is not defined or empty")
	}

	return &kafkaFavoriteConsumerConfig{
		brokers: brokers,
		groupID: groupID,
	}, nil
}

func (c *kafkaFavoriteConsumerConfig) Brokers() []string {
	return c.brokers
}

func (c *kafkaFavoriteConsumerConfig) GroupID() string {
	return c.groupID
}

func (c *kafkaFavoriteConsumerConfig) Config() *sarama.Config {
	config := sarama.NewConfig()

	config.Version = sarama.V2_6_0_0
	config.Consumer.Group.Rebalance.GroupStrategies = []sarama.BalanceStrategy{
		sarama.NewBalanceStrategyRoundRobin(),
	}
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	return config
}
