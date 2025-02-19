package app

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/escoutdoor/kotopes/common/pkg/closer"
	"github.com/escoutdoor/kotopes/common/pkg/db"
	"github.com/escoutdoor/kotopes/common/pkg/db/pg"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"github.com/escoutdoor/kotopes/common/pkg/redis"
	"github.com/escoutdoor/kotopes/common/pkg/transaction"
	pet_implementation "github.com/escoutdoor/kotopes/pet/internal/api/pet_v1"
	"github.com/escoutdoor/kotopes/pet/internal/config"
	"github.com/escoutdoor/kotopes/pet/internal/repository"
	pet_repository "github.com/escoutdoor/kotopes/pet/internal/repository/pet/pg"
	pet_cache "github.com/escoutdoor/kotopes/pet/internal/repository/pet/redis"
	"github.com/escoutdoor/kotopes/pet/internal/service"
	pet_service "github.com/escoutdoor/kotopes/pet/internal/service/pet"
)

type serviceProvider struct {
	grpcServerConfig       config.GRPCConfig
	prometheusServerConfig config.PrometheusConfig

	tracingConfig config.TracingConfig

	pgConfig    config.PGConfig
	redisConfig config.RedisConfig

	validator *protovalidate.Validator

	redisClient   redis.Client
	dbClient      db.Client
	txManager     db.TxManager
	petRepository repository.PetRepository
	petCache      repository.PetCache

	petService service.PetService

	petImpl *pet_implementation.Implementation
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

func (s *serviceProvider) PrometheusConfig() config.PrometheusConfig {
	if s.prometheusServerConfig == nil {
		cfg, err := config.NewPrometheusConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load prometheus server config: %s", err)
		}
		s.prometheusServerConfig = cfg
	}

	return s.prometheusServerConfig
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

func (s *serviceProvider) RedisConfig() config.RedisConfig {
	if s.redisConfig == nil {
		cfg, err := config.NewRedisConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load redis config: %s", err)
		}
		s.redisConfig = cfg
	}

	return s.redisConfig
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

func (s *serviceProvider) RedisClient(ctx context.Context) redis.Client {
	if s.redisClient == nil {
		redisClient, err := redis.NewClient(ctx, s.RedisConfig())
		if err != nil {
			logger.Logger().Fatalf("failed to create redis client: %s", err)
		}

		err = redisClient.Ping(ctx)
		if err != nil {
			logger.Logger().Fatalf("failed to ping redis database: %s", err)
		}

		s.redisClient = redisClient
		closer.Add(redisClient.Close)
	}

	return s.redisClient
}

func (s *serviceProvider) DBClient(ctx context.Context) db.Client {
	if s.dbClient == nil {
		dbClient, err := pg.NewClient(ctx, s.PGConfig().DSN())
		if err != nil {
			logger.Logger().Fatalf("failed to create connection to postgres database: %s", err)
		}

		err = dbClient.DB().Ping(ctx)
		if err != nil {
			logger.Logger().Fatalf("failed to ping postgres database: %s", err)
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

func (s *serviceProvider) PetRepository(ctx context.Context) repository.PetRepository {
	if s.petRepository == nil {
		s.petRepository = pet_repository.New(s.DBClient(ctx))
	}

	return s.petRepository
}

func (s *serviceProvider) PetCache(ctx context.Context) repository.PetCache {
	if s.petCache == nil {
		s.petCache = pet_cache.New(s.RedisClient(ctx))
	}

	return s.petCache
}

func (s *serviceProvider) PetService(ctx context.Context) service.PetService {
	if s.petService == nil {
		s.petService = pet_service.NewService(
			s.PetRepository(ctx),
			s.PetCache(ctx),
			s.TxManager(ctx),
			s.RedisConfig().TTL(),
		)
	}

	return s.petService
}

func (s *serviceProvider) PetImplementation(ctx context.Context) *pet_implementation.Implementation {
	if s.petImpl == nil {
		s.petImpl = pet_implementation.New(s.PetService(ctx))
	}

	return s.petImpl
}
