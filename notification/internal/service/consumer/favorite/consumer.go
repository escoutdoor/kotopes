package favorite

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/escoutdoor/kotopes/common/pkg/kafka"
	grpc_client "github.com/escoutdoor/kotopes/notification/internal/client/grpc"
	"github.com/escoutdoor/kotopes/notification/internal/client/mail"
	"github.com/escoutdoor/kotopes/notification/internal/config"
)

type service struct {
	consumer   kafka.Consumer[sarama.ConsumerMessage]
	config     config.FavoriteServiceConfig
	userClient grpc_client.UserClient
	mailClient mail.Client
}

func New(
	consumer kafka.Consumer[sarama.ConsumerMessage],
	config config.FavoriteServiceConfig,
	userClient grpc_client.UserClient,
	mailClient mail.Client,
) *service {
	return &service{
		consumer:   consumer,
		config:     config,
		userClient: userClient,
		mailClient: mailClient,
	}
}

func (svc *service) RunConsumer(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case err := <-svc.run(ctx):
			if err != nil {
				return err
			}
		}
	}
}

func (svc *service) run(ctx context.Context) <-chan error {
	ch := make(chan error)

	go func() {
		defer close(ch)
		ch <- svc.consumer.Consume(ctx, svc.config.Topic(), svc.FavoriteHandler)
	}()

	return ch
}
