package app

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	pb "github.com/escoutdoor/kotopes/common/api/favorite/v1"
	"github.com/escoutdoor/kotopes/common/pkg/closer"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"github.com/escoutdoor/kotopes/favorite/internal/config"
	errors_interceptor "github.com/escoutdoor/kotopes/favorite/internal/interceptor/errors"
	logging_interceptor "github.com/escoutdoor/kotopes/favorite/internal/interceptor/logging"
	metrics_interceptor "github.com/escoutdoor/kotopes/favorite/internal/interceptor/metrics"
	validation_interceptor "github.com/escoutdoor/kotopes/favorite/internal/interceptor/validation"
	"github.com/escoutdoor/kotopes/favorite/internal/metrics"
	"github.com/escoutdoor/kotopes/favorite/internal/tracing"
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
		a.initServiceProvider,
		a.initConfig,
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

func (a *App) initServiceProvider(_ context.Context) error {
	a.serviceProvider = newServiceProvider()
	return nil
}

func (a *App) initConfig(_ context.Context) error {
	err := config.Load(".env")
	if err != nil {
		return errwrap.Wrap("failed to init config", err)
	}

	return nil
}

func (a *App) initMetrics(_ context.Context) error {
	err := metrics.Init()
	if err != nil {
		return errwrap.Wrap("failed to init prometheus metrics", err)
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

func (a *App) initGRPCServer(ctx context.Context) error {
	grpcServer := grpc.NewServer(
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
		grpc.ChainUnaryInterceptor(
			validation_interceptor.Unary(a.serviceProvider.Validator()),
			metrics_interceptor.Unary(),
			logging_interceptor.Unary(),
			errors_interceptor.Unary(),
		),
	)
	reflection.Register(grpcServer)

	pb.RegisterFavoriteV1Server(grpcServer, a.serviceProvider.FavoriteImplementation(ctx))

	a.grpcServer = grpcServer
	return nil
}

func (a *App) initPrometheusServer(ctx context.Context) error {
	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.Handler())

	a.prometheusServer = &http.Server{
		Addr:              a.serviceProvider.PrometheusConfig().Addr(),
		Handler:           mux,
		ReadHeaderTimeout: time.Second * 5,
	}

	return nil
}

func (a *App) runGRPCServer() error {
	logger.Logger().Info("running grpc server: ", a.serviceProvider.GRPCServerConfig().Addr())
	ln, err := net.Listen("tcp", a.serviceProvider.GRPCServerConfig().Addr())
	if err != nil {
		return errwrap.Wrap("listen tcp", err)
	}

	err = a.grpcServer.Serve(ln)
	if err != nil {
		return errwrap.Wrap("serve", err)
	}
	return nil
}

func (a *App) runPrometheusServer() error {
	logger.Logger().Info("running prometheus server: ", a.serviceProvider.PrometheusConfig().Addr())

	err := a.prometheusServer.ListenAndServe()
	if err != nil {
		return errwrap.Wrap("listen and serve", err)
	}

	return nil
}
