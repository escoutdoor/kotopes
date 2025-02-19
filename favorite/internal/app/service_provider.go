package app

import (
	"context"

	"github.com/IBM/sarama"
	"github.com/bufbuild/protovalidate-go"
	petpb "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/common/pkg/closer"
	"github.com/escoutdoor/kotopes/common/pkg/db"
	"github.com/escoutdoor/kotopes/common/pkg/db/pg"
	"github.com/escoutdoor/kotopes/common/pkg/kafka"
	"github.com/escoutdoor/kotopes/common/pkg/kafka/producer"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"github.com/escoutdoor/kotopes/common/pkg/transaction"
	favorite_implementation "github.com/escoutdoor/kotopes/favorite/internal/api/favorite_v1"
	grpc_client "github.com/escoutdoor/kotopes/favorite/internal/client/grpc"
	pet_client "github.com/escoutdoor/kotopes/favorite/internal/client/grpc/pet"
	kafka_client "github.com/escoutdoor/kotopes/favorite/internal/client/kafka"
	notification_client "github.com/escoutdoor/kotopes/favorite/internal/client/kafka/notification"
	"github.com/escoutdoor/kotopes/favorite/internal/config"
	"github.com/escoutdoor/kotopes/favorite/internal/repository"
	favorite_repository "github.com/escoutdoor/kotopes/favorite/internal/repository/favorite"
	"github.com/escoutdoor/kotopes/favorite/internal/service"
	favorite_service "github.com/escoutdoor/kotopes/favorite/internal/service/favorite"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serviceProvider struct {
	grpcServerConfig config.GRPCConfig
	pgConfig         config.PGConfig

	petServiceConfig config.PetServiceConfig

	prometheusConfig config.PrometheusConfig
	tracingConfig    config.TracingConfig

	notificationServiceConfig config.NotificationServiceConfig
	kafkaProducerConfig       config.KafkaProducerConfig

	validator *protovalidate.Validator

	syncProducer  sarama.SyncProducer
	kafkaProducer kafka.Producer[sarama.ProducerMessage]

	notificationClient kafka_client.NotificationClient

	petGRPCClient petpb.PetV1Client
	petClient     grpc_client.PetClient

	dbClient           db.Client
	txManager          db.TxManager
	favoriteRepository repository.FavoriteRepository

	favoriteService service.FavoriteService

	favoriteImpl *favorite_implementation.Implementation
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) GRPCServerConfig() config.GRPCConfig {
	if s.grpcServerConfig == nil {
		cfg, err := config.NewGRPCConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load grpc server config: %s", err)
		}
		s.grpcServerConfig = cfg
	}

	return s.grpcServerConfig
}

func (s *serviceProvider) PGConfig() config.PGConfig {
	if s.pgConfig == nil {
		cfg, err := config.NewPGConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load postgres config: %s", err)
		}
		s.pgConfig = cfg
	}

	return s.pgConfig
}

func (s *serviceProvider) PetServiceConfig() config.PetServiceConfig {
	if s.petServiceConfig == nil {
		cfg, err := config.NewPetServiceConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load pet service config: %s", err)
		}

		s.petServiceConfig = cfg
	}

	return s.petServiceConfig
}

func (s *serviceProvider) PrometheusConfig() config.PrometheusConfig {
	if s.prometheusConfig == nil {
		cfg, err := config.NewPrometheusConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load prometheus config: %s", err)
		}

		s.prometheusConfig = cfg
	}

	return s.prometheusConfig
}

func (s *serviceProvider) TracingConfig() config.TracingConfig {
	if s.tracingConfig == nil {
		cfg, err := config.NewTracingConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load tracing config: %s", err)
		}

		s.tracingConfig = cfg
	}

	return s.tracingConfig
}

func (s *serviceProvider) NotificationServiceConfig() config.NotificationServiceConfig {
	if s.notificationServiceConfig == nil {
		cfg, err := config.NewNotificationServiceConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load notification service config: %s", err)
		}

		s.notificationServiceConfig = cfg
	}

	return s.notificationServiceConfig
}

func (s *serviceProvider) KafkaProducerConfig() config.KafkaProducerConfig {
	if s.kafkaProducerConfig == nil {
		cfg, err := config.NewKafkaProducerConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load kafka producer config: %s", err)
		}

		s.kafkaProducerConfig = cfg
	}

	return s.kafkaProducerConfig
}

func (s *serviceProvider) Validator() *protovalidate.Validator {
	if s.validator == nil {
		validator, err := protovalidate.New()
		if err != nil {
			logger.Logger().Fatalf("failed to init protovalidate validator: %s", err)
		}

		s.validator = validator
	}

	return s.validator
}

func (s *serviceProvider) SyncProducer() sarama.SyncProducer {
	if s.syncProducer == nil {
		syncProducer, err := sarama.NewSyncProducer(
			s.KafkaProducerConfig().Brokers(),
			s.KafkaProducerConfig().Config(),
		)
		if err != nil {
			logger.Logger().Fatalf("failed to init sync producer: %s", err)
		}

		closer.Add(syncProducer.Close)
		s.syncProducer = syncProducer
	}

	return s.syncProducer
}

func (s *serviceProvider) KafkaProducer() kafka.Producer[sarama.ProducerMessage] {
	if s.kafkaProducer == nil {
		s.kafkaProducer = producer.NewProducer(s.SyncProducer())
	}

	return s.kafkaProducer
}

func (s *serviceProvider) NotificationClient() kafka_client.NotificationClient {
	if s.notificationClient == nil {
		s.notificationClient = notification_client.New(
			s.KafkaProducer(),
			s.NotificationServiceConfig(),
		)
	}

	return s.notificationClient
}

func (s *serviceProvider) PetGRPCClient(ctx context.Context) petpb.PetV1Client {
	if s.petGRPCClient == nil {
		conn, err := grpc.NewClient(
			s.PetServiceConfig().Addr(),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		)
		if err != nil {
			logger.Logger().Fatalf("failed to init pet grpc service client: %s", err)
		}

		client := petpb.NewPetV1Client(conn)
		s.petGRPCClient = client

		closer.Add(conn.Close)
	}

	return s.petGRPCClient
}

func (s *serviceProvider) PetClient(ctx context.Context) grpc_client.PetClient {
	if s.petClient == nil {
		s.petClient = pet_client.New(s.PetGRPCClient(ctx))
	}

	return s.petClient
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		dbClient, err := pg.NewClient(ctx, s.PGConfig().DSN())
		if err != nil {
			logger.Logger().Fatalf("failed to init db client: %s", err)
		}

		err = dbClient.DB().Ping(ctx)
		if err != nil {
			logger.Logger().Fatalf("failed to ping postgres db: %s", err)
		}

		s.dbClient = dbClient
		closer.Add(dbClient.Close)
	}

	return s.dbClient
}

func (s *serviceProvider) TxManager(ctx context.Context) db.TxManager {
	if s.txManager == nil {
		m := transaction.NewTransactionManager(s.DBClient(ctx).DB())
		s.txManager = m
	}

	return s.txManager
}

func (s *serviceProvider) FavoriteRepository(ctx context.Context) repository.FavoriteRepository {
	if s.favoriteRepository == nil {
		s.favoriteRepository = favorite_repository.New(s.DBClient(ctx))
	}

	return s.favoriteRepository
}

func (s *serviceProvider) FavoriteService(ctx context.Context) service.FavoriteService {
	if s.favoriteService == nil {
		s.favoriteService = favorite_service.New(
			s.FavoriteRepository(ctx),
			s.TxManager(ctx),
			s.PetClient(ctx),
			s.NotificationClient(),
		)
	}

	return s.favoriteService
}

func (s *serviceProvider) FavoriteImplementation(ctx context.Context) *favorite_implementation.Implementation {
	if s.favoriteImpl == nil {
		s.favoriteImpl = favorite_implementation.New(s.FavoriteService(ctx))
	}

	return s.favoriteImpl
}
