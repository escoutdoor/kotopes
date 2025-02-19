package producer

import (
	"github.com/IBM/sarama"

	"github.com/escoutdoor/kotopes/common/pkg/kafka"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
)

var _ kafka.Producer[sarama.ProducerMessage] = (*kafkaProducer)(nil)

type kafkaProducer struct {
	producer sarama.SyncProducer
}

// NewProducer creates producer for kafka
func NewProducer(producer sarama.SyncProducer) *kafkaProducer {
	return &kafkaProducer{
		producer: producer,
	}
}

func (p *kafkaProducer) SendMessage(msg *sarama.ProducerMessage) error {
	partition, offset, err := p.producer.SendMessage(msg)
	if err != nil {
		return err
	}

	logger.Logger().Infof("message sent to partition %d with offset %d\n", partition, offset)

	return nil
}

func (p *kafkaProducer) Close() error {
	return p.producer.Close()
}
