package app

import (
	"context"
	"net"
	"net/http"
	"sync"
	"time"

	pb "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/common/pkg/closer"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"github.com/escoutdoor/kotopes/pet/internal/config"
	errors_interceptor "github.com/escoutdoor/kotopes/pet/internal/interceptor/errors"
	logging_interceptor "github.com/escoutdoor/kotopes/pet/internal/interceptor/logging"
	metrics_interceptor "github.com/escoutdoor/kotopes/pet/internal/interceptor/metrics"
	validation_interceptor "github.com/escoutdoor/kotopes/pet/internal/interceptor/validation"
	"github.com/escoutdoor/kotopes/pet/internal/metrics"
	"github.com/escoutdoor/kotopes/pet/internal/tracing"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"

	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	err := a.initDeps(ctx)
	if err != nil {
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

		err := a.runGRPCServer()
		if err != nil {
			logger.Logger().Fatalf("failed to run grpc server: %s\n", err)
		}
	}()

	go func() {
		defer wg.Done()

		err := a.runPrometheusServer()
		if err != nil {
			logger.Logger().Fatalf("failed to run prometheus server: %s\n", err)
		}
	}()

	wg.Wait()
	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	initFuncs := []func(ctx context.Context) error{
		a.initConfig,
		a.initServiceProvider,
		a.initTracing,
		a.initMetrics,
		a.initPrometheusServer,
		a.initGRPCServer,
	}

	for _, fn := range initFuncs {
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
		return errwrap.Wrap("failed to init config", err)
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
			metrics_interceptor.Unary(),
			logging_interceptor.Unary(),
			errors_interceptor.Unary(),
		),
	)
	reflection.Register(grpcServer)
	pb.RegisterPetV1Server(grpcServer, a.serviceProvider.PetImplementation(ctx))

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
		return errwrap.Wrap("failed to init jaeger tracing", err)
	}

	return nil
}

func (a *App) runGRPCServer() error {
	logger.Logger().Info("grpc server is running: ", a.serviceProvider.GRPCServerConfig().Addr())

	ln, err := net.Listen("tcp", a.serviceProvider.GRPCServerConfig().Addr())
	if err != nil {
		return errwrap.Wrap("listen tcp", err)
	}

	err = a.grpcServer.Serve(ln)
	if err != nil {
		return errwrap.Wrap("server grpc server", err)
	}
	return nil
}

func (a *App) runPrometheusServer() error {
	logger.Logger().Info("prometheus server is running: ",
		a.serviceProvider.PrometheusConfig().Addr())

	err := a.prometheusServer.ListenAndServe()
	if err != nil {
		return errwrap.Wrap("list and server http server", err)
	}

	return nil
}
