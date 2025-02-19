package main

import (
	"context"
	"flag"

	"github.com/escoutdoor/kotopes/auth/internal/app"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"go.uber.org/zap/zapcore"
)

var configPath string

func init() {
	flag.StringVar(&configPath, "config", ".env", "path to configuration file")
	flag.Parse()
}

func main() {
	ctx := context.Background()
	logger.SetLevel(zapcore.DebugLevel)

	logger.Info(ctx, "application init start")
	a, err := app.New(ctx, configPath)
	if err != nil {
		logger.Fatal(ctx, err)
	}

	logger.Info(ctx, "running application")
	err = a.Run()
	if err != nil {
		logger.Fatal(ctx, err)
	}
}
