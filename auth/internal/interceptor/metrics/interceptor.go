package metrics

import (
	"context"
	"time"

	"github.com/escoutdoor/kotopes/auth/internal/metrics"
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
		metrics.IncRequestCounter()
		start := time.Now()

		resp, err := handler(ctx, req)
		diff := time.Since(start)
		code := status.Code(err).String()

		if err != nil {
			metrics.IncResponseCounter(code, info.FullMethod)
			metrics.HistogramResponseTimeObserve(code, diff.Seconds())
		} else {
			metrics.IncResponseCounter(code, info.FullMethod)
			metrics.HistogramResponseTimeObserve(code, diff.Seconds())
		}

		return resp, err
	}
}
