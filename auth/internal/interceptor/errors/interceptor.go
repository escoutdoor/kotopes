package errors

import (
	"context"
	stderrs "errors"

	"github.com/escoutdoor/kotopes/auth/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func Unary() grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		resp, err := handler(ctx, req)

		_, ok := status.FromError(err)
		if ok {
			return resp, err
		}

		switch {
		case stderrs.Is(err, model.ErrUserAlreadyExists):
			err = status.Error(codes.AlreadyExists, err.Error())
		case stderrs.Is(err, model.ErrUserNotFound):
			err = status.Error(codes.NotFound, err.Error())
		case stderrs.Is(err, model.ErrInvalidEmailOrPassword):
			err = status.Error(codes.InvalidArgument, err.Error())
		default:
			err = status.Error(codes.Internal, err.Error())
		}

		return resp, err
	}
}
