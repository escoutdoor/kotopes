package validation

import (
	"context"

	"github.com/bufbuild/protovalidate-go"
	"github.com/escoutdoor/kotopes/common/pkg/grpcutil"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"google.golang.org/protobuf/proto"
)

func Unary(validator *protovalidate.Validator) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		msg, ok := req.(proto.Message)
		if !ok {
			return nil, status.Errorf(codes.Internal, "unsupported message type: %T", msg)
		}

		err := validator.Validate(msg)
		if err != nil {
			return nil, grpcutil.ProtoValidationError(err)
		}

		return handler(ctx, req)
	}
}
