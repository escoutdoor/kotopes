package auth_v1

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/converter"
	pb "github.com/escoutdoor/kotopes/common/api/auth/v1"
)

func (i *Implementation) Register(ctx context.Context, req *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	id, err := i.authService.Register(ctx, converter.ToCreateUserFromPb(req))
	if err != nil {
		return nil, err
	}

	return &pb.RegisterResponse{
		Id: id,
	}, nil
}
