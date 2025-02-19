package logging

import (
	"context"
	"time"

	"github.com/escoutdoor/kotopes/common/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		logCtx := logger.ToContext(context.Background(),
			logger.FromContext(ctx).With(
				"operation", info.FullMethod,
				"component", "interceptor",
			),
		)
		start := time.Now()

		logger.Debug(logCtx, "receive request")
		resp, err := handler(ctx, req)
		logger.Debug(
			logCtx,
			"handle request",
			zap.Any("req", req),
			zap.Any("resp", resp), // may contain sensitive data
			zap.Int64("duration milliseconds", time.Since(start).Milliseconds()),
		)

		if err != nil {
			logger.Error(
				logCtx,
				"handle error",
				zap.String("status", status.Code(err).String()),
				zap.Any("req", req),
				zap.String("error", err.Error()),
			)
		}

		return resp, err
	}

}
