package auth_v1

import (
	"context"

	pb "github.com/escoutdoor/kotopes/common/api/auth/v1"
)

func (i *Implementation) Validate(ctx context.Context, req *pb.ValidateRequest) (*pb.ValidateResponse, error) {
	data, err := i.authService.Validate(ctx, req.GetAccessToken())
	if err != nil {
		return nil, err
	}

	return &pb.ValidateResponse{
		Id:   data.ID,
		Role: data.Role,
	}, nil
}
