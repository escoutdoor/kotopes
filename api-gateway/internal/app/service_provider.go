package app

import (
	grpc_client "github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc"
	access_client "github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc/access"
	auth_client "github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc/auth"
	favorite_client "github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc/favorite"
	pet_client "github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc/pet"
	user_client "github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc/user"
	"github.com/escoutdoor/kotopes/api-gateway/internal/config"
	access_v1 "github.com/escoutdoor/kotopes/common/api/access/v1"
	auth_v1 "github.com/escoutdoor/kotopes/common/api/auth/v1"
	favorite_v1 "github.com/escoutdoor/kotopes/common/api/favorite/v1"
	pet_v1 "github.com/escoutdoor/kotopes/common/api/pet/v1"
	user_v1 "github.com/escoutdoor/kotopes/common/api/user/v1"
	"github.com/escoutdoor/kotopes/common/pkg/closer"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type serviceProvider struct {
	httpServerConfig       config.HTTPServerConfig
	prometheusServerConfig config.PrometheusServerConfig

	tracingConfig config.TracingConfig

	petClientConfig      config.GRPCServiceClientConfig
	authClientConfig     config.GRPCServiceClientConfig
	favoriteClientConfig config.GRPCServiceClientConfig

	authServiceConnection *grpc.ClientConn

	authGRPCClient     auth_v1.AuthV1Client
	accessGRPCClient   access_v1.AccessV1Client
	userGRPCClient     user_v1.UserV1Client
	petGRPCClient      pet_v1.PetV1Client
	favoriteGRPCClient favorite_v1.FavoriteV1Client

	authClient     grpc_client.AuthClient
	accessClient   grpc_client.AccessClient
	userClient     grpc_client.UserClient
	petClient      grpc_client.PetClient
	favoriteClient grpc_client.FavoriteClient
}

func newServiceProvider() *serviceProvider {
	return &serviceProvider{}
}

func (s *serviceProvider) HTTPServerConfig() config.HTTPServerConfig {
	if s.httpServerConfig == nil {
		cfg, err := config.NewHTTPServerConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load http server config: %s", err)
		}

		s.httpServerConfig = cfg
	}

	return s.httpServerConfig
}

func (s *serviceProvider) PrometheusServerConfig() config.PrometheusServerConfig {
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

func (s *serviceProvider) PetClientConfig() config.GRPCServiceClientConfig {
	if s.petClientConfig == nil {
		cfg, err := config.NewPetClientConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load pet service client config: %s", err)
		}

		s.petClientConfig = cfg
	}

	return s.petClientConfig
}

func (s *serviceProvider) AuthClientConfig() config.GRPCServiceClientConfig {
	if s.authClientConfig == nil {
		cfg, err := config.NewAuthClientConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load auth service client config: %s", err)
		}

		s.authClientConfig = cfg
	}

	return s.authClientConfig
}

func (s *serviceProvider) FavoriteClientConfig() config.GRPCServiceClientConfig {
	if s.favoriteClientConfig == nil {
		cfg, err := config.NewFavoriteClientConfig()
		if err != nil {
			logger.Logger().Fatalf("failed to load favorite service client config: %s", err)
		}

		s.favoriteClientConfig = cfg
	}

	return s.favoriteClientConfig
}

func (s *serviceProvider) AuthServiceConnection() *grpc.ClientConn {
	if s.authServiceConnection == nil {
		conn, err := grpc.NewClient(
			s.AuthClientConfig().Addr(),
			// TODO: implement secured dial
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		)
		if err != nil {
			logger.Logger().Fatalf("failed to connect to auth grpc service: %s", err)
		}

		s.authServiceConnection = conn
		closer.Add(conn.Close)
	}

	return s.authServiceConnection
}

func (s *serviceProvider) AuthGRPCClient() auth_v1.AuthV1Client {
	if s.authGRPCClient == nil {
		s.authGRPCClient = auth_v1.NewAuthV1Client(s.AuthServiceConnection())
	}

	return s.authGRPCClient
}

func (s *serviceProvider) AccessGRPCClient() access_v1.AccessV1Client {
	if s.accessGRPCClient == nil {
		s.accessGRPCClient = access_v1.NewAccessV1Client(s.AuthServiceConnection())
	}

	return s.accessGRPCClient
}

func (s *serviceProvider) UserGRPCClient() user_v1.UserV1Client {
	if s.userGRPCClient == nil {
		s.userGRPCClient = user_v1.NewUserV1Client(s.AuthServiceConnection())
	}

	return s.userGRPCClient
}

func (s *serviceProvider) PetGRPCClient() pet_v1.PetV1Client {
	if s.petGRPCClient == nil {
		conn, err := grpc.NewClient(
			s.PetClientConfig().Addr(),
			// TODO: implement secured dial
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		)
		if err != nil {
			logger.Logger().Fatalf("failed to connect to pet grpc service: %s", err)
		}

		s.petGRPCClient = pet_v1.NewPetV1Client(conn)
		closer.Add(conn.Close)
	}

	return s.petGRPCClient
}

func (s *serviceProvider) FavoriteGRPCClient() favorite_v1.FavoriteV1Client {
	if s.favoriteGRPCClient == nil {
		conn, err := grpc.NewClient(
			s.FavoriteClientConfig().Addr(),
			// TODO: implement secured dial
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		)
		if err != nil {
			logger.Logger().Fatalf("failed to connect to favorite grpc service: %s", err)
		}

		s.favoriteGRPCClient = favorite_v1.NewFavoriteV1Client(conn)
		closer.Add(conn.Close)
	}

	return s.favoriteGRPCClient
}

func (s *serviceProvider) AuthClient() grpc_client.AuthClient {
	if s.authClient == nil {
		s.authClient = auth_client.New(s.AuthGRPCClient())
	}

	return s.authClient
}

func (s *serviceProvider) AccessClient() grpc_client.AccessClient {
	if s.accessClient == nil {
		s.accessClient = access_client.New(s.AccessGRPCClient())
	}

	return s.accessClient
}

func (s *serviceProvider) UserClient() grpc_client.UserClient {
	if s.userClient == nil {
		s.userClient = user_client.New(s.UserGRPCClient())
	}

	return s.userClient
}

func (s *serviceProvider) PetClient() grpc_client.PetClient {
	if s.petClient == nil {
		s.petClient = pet_client.New(s.PetGRPCClient())
	}

	return s.petClient
}

func (s *serviceProvider) FavoriteClient() grpc_client.FavoriteClient {
	if s.favoriteClient == nil {
		s.favoriteClient = favorite_client.New(s.FavoriteGRPCClient())
	}

	return s.favoriteClient
}
