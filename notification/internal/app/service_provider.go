package app

import (
	"github.com/IBM/sarama"
	userpb "github.com/escoutdoor/kotopes/common/api/user/v1"
	"github.com/escoutdoor/kotopes/common/pkg/closer"
	"github.com/escoutdoor/kotopes/common/pkg/kafka"
	grpc_client "github.com/escoutdoor/kotopes/notification/internal/client/grpc"
	user_client "github.com/escoutdoor/kotopes/notification/internal/client/grpc/user"
	"github.com/escoutdoor/kotopes/notification/internal/client/mail"
	go_mail_client "github.com/escoutdoor/kotopes/notification/internal/client/mail/go-mail"
	"github.com/escoutdoor/kotopes/notification/internal/config"
	"github.com/escoutdoor/kotopes/notification/internal/service"
	favorite_service "github.com/escoutdoor/kotopes/notification/internal/service/consumer/favorite"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"github.com/escoutdoor/kotopes/common/pkg/kafka/consumer"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
)

type serviceProvider struct {
	mailConfig config.MailConfig

	kafkaFavoriteConsumerConfig config.KafkaConsumerConfig

	userServiceConfig     config.UserServiceConfig
	favoriteServiceConfig config.FavoriteServiceConfig

	mailClient mail.Client

	userGRPCClient userpb.UserV1Client
	userClient     grpc_client.UserClient

	favoriteConsumerService service.ConsumerService

	kafkaFavoriteConsumerGroup        sarama.ConsumerGroup
	kafkaFavoriteConsumerGroupHandler *consumer.GroupHandler
	kafkaFavoriteConsumer             kafka.Consumer[sarama.ConsumerMessage]
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) MailConfig() config.MailConfig {
	if s.mailConfig == nil {
		cfg, err := config.NewMailConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load mail config: %s", err)
		}

		s.mailConfig = cfg
	}

	return s.mailConfig
}

func (s *serviceProvider) KafkaFavoriteConsumerConfig() config.KafkaConsumerConfig {
	if s.kafkaFavoriteConsumerConfig == nil {
		cfg, err := config.NewKafkaFavoriteConsumerConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load kafka favorite consumer config: %s", err)
		}

		s.kafkaFavoriteConsumerConfig = cfg
	}

	return s.kafkaFavoriteConsumerConfig
}

func (s *serviceProvider) UserServiceConfig() config.UserServiceConfig {
	if s.userServiceConfig == nil {
		cfg, err := config.NewUserServiceConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load user service config: %s", err)
		}

		s.userServiceConfig = cfg
	}

	return s.userServiceConfig
}

func (s *serviceProvider) FavoriteServiceConfig() config.FavoriteServiceConfig {
	if s.favoriteServiceConfig == nil {
		cfg, err := config.NewFavoriteServiceConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load favorite service config: %s", err)
		}

		s.favoriteServiceConfig = cfg
	}

	return s.favoriteServiceConfig
}

func (s *serviceProvider) MailClient() mail.Client {
	if s.mailClient == nil {
		cl, err := go_mail_client.NewClient(s.MailConfig())
		if err != nil {
			logger.Logger().Fatalf("failed to init go-mail client: %s", err)
		}

		s.mailClient = cl
		closer.Add(cl.Close)
	}

	return s.mailClient
}

func (s *serviceProvider) UserGRPCClient() userpb.UserV1Client {
	if s.userGRPCClient == nil {
		conn, err := grpc.NewClient(
			s.UserServiceConfig().Addr(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			logger.Logger().Fatalf("failed to connect to user grpc client: %s", err)
		}

		s.userGRPCClient = userpb.NewUserV1Client(conn)
		closer.Add(conn.Close)
	}

	return s.userGRPCClient
}

func (s *serviceProvider) UserClient() grpc_client.UserClient {
	if s.userClient == nil {
		s.userClient = user_client.New(s.UserGRPCClient())
	}

	return s.userClient
}

func (s *serviceProvider) FavoriteConsumerService() service.ConsumerService {
	if s.favoriteConsumerService == nil {
		s.favoriteConsumerService = favorite_service.New(
			s.KafkaFavoriteConsumer(),
			s.FavoriteServiceConfig(),
			s.UserClient(),
			s.MailClient(),
		)
	}

	return s.favoriteConsumerService
}

func (s *serviceProvider) KafkaFavoriteConsumerGroup() sarama.ConsumerGroup {
	if s.kafkaFavoriteConsumerGroup == nil {
		cgroup, err := sarama.NewConsumerGroup(
			s.KafkaFavoriteConsumerConfig().Brokers(),
			s.KafkaFavoriteConsumerConfig().GroupID(),
			s.KafkaFavoriteConsumerConfig().Config(),
		)
		if err != nil {
			logger.Logger().Fatalf("failed to init kafka favorite consumer group: %s", err)
		}

		s.kafkaFavoriteConsumerGroup = cgroup
	}

	return s.kafkaFavoriteConsumerGroup
}

func (s *serviceProvider) KafkaFavoriteConsumerGroupHandler() *consumer.GroupHandler {
	if s.kafkaFavoriteConsumerGroupHandler == nil {
		s.kafkaFavoriteConsumerGroupHandler = consumer.NewGroupHandler()
	}

	return s.kafkaFavoriteConsumerGroupHandler
}

func (s *serviceProvider) KafkaFavoriteConsumer() kafka.Consumer[sarama.ConsumerMessage] {
	if s.kafkaFavoriteConsumer == nil {
		s.kafkaFavoriteConsumer = consumer.NewConsumer(
			s.KafkaFavoriteConsumerGroup(),
			s.KafkaFavoriteConsumerGroupHandler(),
		)
	}

	return s.kafkaFavoriteConsumer
}
