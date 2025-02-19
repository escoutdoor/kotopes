package app

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"sync"
	"time"

	"github.com/escoutdoor/kotopes/auth/internal/config"
	errors_interceptor "github.com/escoutdoor/kotopes/auth/internal/interceptor/errors"
	logging_interceptor "github.com/escoutdoor/kotopes/auth/internal/interceptor/logging"
	metrics_interceptor "github.com/escoutdoor/kotopes/auth/internal/interceptor/metrics"
	validation_interceptor "github.com/escoutdoor/kotopes/auth/internal/interceptor/validation"
	"github.com/escoutdoor/kotopes/auth/internal/metrics"
	"github.com/escoutdoor/kotopes/auth/internal/tracing"
	accesspb "github.com/escoutdoor/kotopes/common/api/access/v1"
	authpb "github.com/escoutdoor/kotopes/common/api/auth/v1"
	userpb "github.com/escoutdoor/kotopes/common/api/user/v1"
	"github.com/escoutdoor/kotopes/common/pkg/closer"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type App struct {
	grpcServer       *grpc.Server
	prometheusServer *http.Server

	serviceProvider *serviceProvider
	configPath      string
}

func New(ctx context.Context, configPath string) (*App, error) {
	a := &App{configPath: configPath}
	if err := a.initDeps(ctx); err != nil {
		return nil, err
	}

	return a, nil
}

func (a *App) Run() error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		if err := a.runGRPCServer(); err != nil {
			logger.Logger().Fatalf("failed to run grpc server: %s", err)
		}
	}()

	go func() {
		defer wg.Done()
		if err := a.runPrometheusServer(); err != nil {
			logger.Logger().Fatalf("failed to run prometheus server: %s", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initMetrics,
		a.initTracing,
		a.initGRPCServer,
		a.initPrometheusServer,
	}

	for _, fn := range deps {
		err := fn(ctx)
		if err != nil {
			return err
		}
	}

	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(a.configPath)
	if err != nil {
		return fmt.Errorf("failed to load config: %s", err)
	}
	return nil
}

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initGRPCServer(ctx context.Context) error {
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			validation_interceptor.Unary(a.serviceProvider.Validator()),
			logging_interceptor.Unary(),
			metrics_interceptor.Unary(),
			errors_interceptor.Unary(),
		),
	)
	reflection.Register(grpcServer)

	authpb.RegisterAuthV1Server(grpcServer, a.serviceProvider.AuthImplementation(ctx))
	userpb.RegisterUserV1Server(grpcServer, a.serviceProvider.UserImplementation(ctx))
	accesspb.RegisterAccessV1Server(grpcServer, a.serviceProvider.AccessImplementation(ctx))

	a.grpcServer = grpcServer
	closer.Add(func() error {
		grpcServer.GracefulStop()
		return nil
	})
	return nil
}

func (a *App) initPrometheusServer(_ context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheusServer = &http.Server{
		Addr:              a.serviceProvider.PrometheusConfig().Addr(),
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 5,
	}
	return nil
}

func (a *App) initMetrics(ctx context.Context) error {
	err := metrics.Init(ctx)
	if err != nil {
		return errwrap.Wrap("failed to init metrics", err)
	}
	return nil
}

func (a *App) initTracing(ctx context.Context) error {
	err := tracing.Init(ctx, a.serviceProvider.TracingConfig())
	if err != nil {
		return errwrap.Wrap("failed to init tracing", err)
	}
	return nil
}

func (a *App) runGRPCServer() error {
	logger.Logger().Info("grpc server is running: ", a.serviceProvider.GRPCServerConfig().Addr())

	ln, err := net.Listen("tcp", a.serviceProvider.GRPCServerConfig().Addr())
	if err != nil {
		return fmt.Errorf("net listen error: %s", err)
	}

	err = a.grpcServer.Serve(ln)
	if err != nil {
		return fmt.Errorf("server server error: %s", err)
	}
	return nil
}

func (a *App) runPrometheusServer() error {
	logger.Logger().Info("prometheus server is running: ", a.serviceProvider.GRPCServerConfig().Addr())

	err := a.prometheusServer.ListenAndServe()
	if err != nil {
		return errwrap.Wrap("listen and server http server", err)
	}
	return nil
}
