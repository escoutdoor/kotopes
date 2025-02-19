package app

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	access_implementation "github.com/escoutdoor/kotopes/auth/internal/api/access_v1"
	auth_implementation "github.com/escoutdoor/kotopes/auth/internal/api/auth_v1"
	user_implementation "github.com/escoutdoor/kotopes/auth/internal/api/user_v1"
	"github.com/escoutdoor/kotopes/auth/internal/config"
	"github.com/escoutdoor/kotopes/auth/internal/repository"
	user_repository "github.com/escoutdoor/kotopes/auth/internal/repository/user"
	"github.com/escoutdoor/kotopes/auth/internal/service"
	access_service "github.com/escoutdoor/kotopes/auth/internal/service/access"
	auth_service "github.com/escoutdoor/kotopes/auth/internal/service/auth"
	user_service "github.com/escoutdoor/kotopes/auth/internal/service/user"
	"github.com/escoutdoor/kotopes/auth/internal/utils/hasher"
	"github.com/escoutdoor/kotopes/auth/internal/utils/policy"
	"github.com/escoutdoor/kotopes/auth/internal/utils/token"
	"github.com/escoutdoor/kotopes/common/pkg/closer"
	"github.com/escoutdoor/kotopes/common/pkg/db"
	"github.com/escoutdoor/kotopes/common/pkg/db/pg"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"github.com/escoutdoor/kotopes/common/pkg/redis"
	"github.com/escoutdoor/kotopes/common/pkg/transaction"
)

type serviceProvider struct {
	grpcServerConfig config.GRPCConfig

	pgConfig    config.PGConfig
	redisConfig config.RedisConfig

	accessTokenConfig  config.TokenConfig
	refreshTokenConfig config.TokenConfig

	tracingConfig    config.TracingConfig
	prometheusConfig config.PrometheusConfig

	redisClient    redis.Client
	dbClient       db.Client
	txManager      db.TxManager
	userRepository repository.UserRepository

	userService   service.UserService
	authService   service.AuthService
	accessService service.AccessService

	accessTokenProvider  token.Provider
	refreshTokenProvider token.Provider
	passwordHasher       hasher.Hasher

	policy *policy.Policy

	validator *protovalidate.Validator

	userImpl   *user_implementation.Implementation
	authImpl   *auth_implementation.Implementation
	accessImpl *access_implementation.Implementation
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
			logger.Logger().Fatalf("failed to load grpc server config: %s", err)
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

func (s *serviceProvider) AccessTokenConfig() config.TokenConfig {
	if s.accessTokenConfig == nil {
		cfg, err := config.NewAccessTokenConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load access token config: %s", err)
		}
		s.accessTokenConfig = cfg
	}

	return s.accessTokenConfig
}

func (s *serviceProvider) RefreshTokenConfig() config.TokenConfig {
	if s.refreshTokenConfig == nil {
		cfg, err := config.NewRefreshTokenConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load refresh token config: %s", err)
		}
		s.refreshTokenConfig = cfg
	}

	return s.refreshTokenConfig
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

func (s *serviceProvider) RedisClient(ctx context.Context) redis.Client {
	if s.redisClient == nil {
		redisClient, err := redis.NewClient(ctx, s.RedisConfig())
		if err != nil {
			logger.Logger().Fatalf("failed to init redis client: %s", err)
		}

		err = redisClient.Ping(ctx)
		if err != nil {
			logger.Logger().Fatalf("failed to ping redis db: %s", err)
		}

		closer.Add(redisClient.Close)
		s.redisClient = redisClient
	}

	return s.redisClient
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

		closer.Add(dbClient.Close)
		s.dbClient = dbClient
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

func (s *serviceProvider) UserRepository(ctx context.Context) repository.UserRepository {
	if s.userRepository == nil {
		s.userRepository = user_repository.New(s.DBClient(ctx))
	}

	return s.userRepository
}

func (s *serviceProvider) UserService(ctx context.Context) service.UserService {
	if s.userService == nil {
		s.userService = user_service.New(s.UserRepository(ctx))
	}

	return s.userService
}

func (s *serviceProvider) AuthService(ctx context.Context) service.AuthService {
	if s.authService == nil {
		s.authService = auth_service.New(
			s.UserRepository(ctx),
			s.AccessTokenProvider(),
			s.RefreshTokenProvider(),
			s.PasswordHasher(),
		)
	}

	return s.authService
}

func (s *serviceProvider) AccessService() service.AccessService {
	if s.accessService == nil {
		s.accessService = access_service.New(
			s.Policy(),
		)
	}

	return s.accessService
}

func (s *serviceProvider) AccessTokenProvider() token.Provider {
	if s.accessTokenProvider == nil {
		s.accessTokenProvider = token.NewTokenProvider(s.AccessTokenConfig())
	}

	return s.accessTokenProvider
}

func (s *serviceProvider) RefreshTokenProvider() token.Provider {
	if s.refreshTokenProvider == nil {
		s.refreshTokenProvider = token.NewTokenProvider(s.RefreshTokenConfig())
	}

	return s.refreshTokenProvider
}

func (s *serviceProvider) PasswordHasher() hasher.Hasher {
	if s.passwordHasher == nil {
		s.passwordHasher = hasher.NewBcryptHasher()
	}

	return s.passwordHasher
}

func (s *serviceProvider) Policy() *policy.Policy {
	if s.policy == nil {
		p, err := policy.New("internal/utils/policy/rbac.rego")
		if err != nil {
			logger.Logger().Fatalf("failed to init policy rules: %s", err)
		}

		s.policy = p
	}

	return s.policy
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

func (s *serviceProvider) AuthImplementation(ctx context.Context) *auth_implementation.Implementation {
	if s.authImpl == nil {
		s.authImpl = auth_implementation.New(s.AuthService(ctx))
	}

	return s.authImpl
}

func (s *serviceProvider) UserImplementation(ctx context.Context) *user_implementation.Implementation {
	if s.userImpl == nil {
		s.userImpl = user_implementation.New(s.UserService(ctx))
	}

	return s.userImpl
}

func (s *serviceProvider) AccessImplementation(ctx context.Context) *access_implementation.Implementation {
	if s.accessImpl == nil {
		s.accessImpl = access_implementation.New(s.AccessService())
	}

	return s.accessImpl
}
