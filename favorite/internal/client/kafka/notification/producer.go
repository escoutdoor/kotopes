package notification

import (
	"github.com/IBM/sarama"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/common/pkg/kafka"
	"github.com/escoutdoor/kotopes/favorite/internal/client/kafka/notification/converter"
	"github.com/escoutdoor/kotopes/favorite/internal/config"
	"github.com/escoutdoor/kotopes/favorite/internal/model"
	"golang.org/x/net/context"
)

type notificationClient struct {
	producer kafka.Producer[sarama.ProducerMessage]
	config   config.NotificationServiceConfig
}

func New(
	producer kafka.Producer[sarama.ProducerMessage],
	config config.NotificationServiceConfig,
) *notificationClient {
	return &notificationClient{
		producer: producer,
		config:   config,
	}
}

func (cl *notificationClient) Send(ctx context.Context, notification *model.Notification) error {
	const op = "notification_client.Send"

	msg, err := converter.ToKafkaMsgFromNotification(notification, cl.config.Topic())
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	err = cl.producer.SendMessage(msg)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}
