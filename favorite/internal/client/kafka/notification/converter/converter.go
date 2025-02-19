package converter

import (
	"encoding/json"
	"fmt"

	"github.com/IBM/sarama"
	"github.com/escoutdoor/kotopes/favorite/internal/model"
)

func ToKafkaMsgFromNotification(v *model.Notification, topic string) (*sarama.ProducerMessage, error) {
	encv, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("marshal struct error: %s", err)
	}

	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(encv),
	}
	return msg, nil
}
