package app

import (
	"context"
	"sync"

	"github.com/escoutdoor/kotopes/common/pkg/closer"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"github.com/escoutdoor/kotopes/notification/internal/config"
)

type App struct {
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

func (a *App) Run(ctx context.Context) error {
	defer func() {
		closer.CloseAll()
		closer.Wait()
	}()
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		logger.Info(ctx,
			"running kafka favorite consumer: ",
			a.serviceProvider.KafkaFavoriteConsumerConfig().Brokers(),
		)
		err := a.serviceProvider.FavoriteConsumerService().RunConsumer(ctx)
		if err != nil {
			logger.Fatalf(ctx, "failed to run kafka favorite consumer: %s", err)
		}
	}()

	wg.Wait()

	return nil
}

func (a *App) initDeps(ctx context.Context) error {
	deps := []func(context.Context) error{
		a.initServiceProvider,
		a.initConfig,
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
	err := config.Load(a.configPath)
	if err != nil {
		return errwrap.Wrap("failed to init config", err)
	}

	return nil
}
