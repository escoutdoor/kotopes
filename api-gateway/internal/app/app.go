package app

import (
	"context"
	"net/http"
	"sync"
	"time"

	"github.com/escoutdoor/kotopes/api-gateway/internal/config"
	auth_v1 "github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/auth/v1"
	favorite_v1 "github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/favorite/v1"
	pet_v1 "github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/pet/v1"
	user_v1 "github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/user/v1"
	"github.com/escoutdoor/kotopes/api-gateway/internal/metrics"
	"github.com/escoutdoor/kotopes/api-gateway/internal/middleware"
	"github.com/escoutdoor/kotopes/api-gateway/internal/tracing"
	"github.com/escoutdoor/kotopes/common/pkg/closer"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"github.com/go-chi/chi/v5"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/cors"
	"go.opentelemetry.io/otel"
)

type App struct {
	httpServer *http.Server

	serviceProvider *serviceProvider
	configPath      string
}

func New(ctx context.Context, configPath string) (*App, error) {
	app := &App{configPath: configPath}

	err := app.initDeps(ctx)
	if err != nil {
		return nil, err
	}

	return app, nil
}

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()

	wg := &sync.WaitGroup{}
	wg.Add(1)

	go func() {
		defer wg.Done()
		if err := a.runHttpServer(ctx); err != nil {
			logger.Fatalf(ctx, "failed to run http server: %s", err)
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
		a.initHttpServer,
	}

	for _, fn := range deps {
		err := fn(ctx)
		if err != nil {
			return err
		}
	}
	return nil
}

func (a *App) initConfig(ctx context.Context) error {
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

func (a *App) initMetrics(_ context.Context) error {
	err := metrics.Init()
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

func (a *App) initHttpServer(_ context.Context) error {
	router := chi.NewRouter()
	router.Use(middleware.TracingMiddleware(otel.Tracer("")))
	router.Handle("/metrics", promhttp.Handler())

	authMiddleware := middleware.AuthMiddleware(
		a.serviceProvider.AuthClient(),
		a.serviceProvider.AccessClient(),
	)

	pet_v1.Register(
		router,
		a.serviceProvider.PetClient(),
		authMiddleware,
	)
	auth_v1.Register(
		router,
		a.serviceProvider.AuthClient(),
	)
	user_v1.Register(
		router,
		a.serviceProvider.UserClient(),
		authMiddleware,
	)
	favorite_v1.Register(
		router,
		a.serviceProvider.FavoriteClient(),
		authMiddleware,
	)

	corsMiddleware := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Authorization"},
		AllowCredentials: true,
	})

	server := &http.Server{
		Addr:              a.serviceProvider.HTTPServerConfig().Addr(),
		Handler:           corsMiddleware.Handler(router),
		ReadHeaderTimeout: time.Second * 5,
	}

	a.httpServer = server
	closer.Add(server.Close)
	return nil
}

func (a *App) runHttpServer(ctx context.Context) error {
	logger.Info(ctx, "running http server: ", a.serviceProvider.HTTPServerConfig().Addr())

	err := a.httpServer.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}
