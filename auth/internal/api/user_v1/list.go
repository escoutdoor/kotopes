package user_v1

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/converter"
	pb "github.com/escoutdoor/kotopes/common/api/user/v1"
)

func (i *Implementation) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	users, err := i.userService.List(ctx, req.GetUserIds())
	if err != nil {
		return nil, err
	}

	return &pb.ListResponse{
		Users: converter.ToPbFromUsers(users),
	}, nil
}
