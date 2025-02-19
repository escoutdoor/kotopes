package user_v1

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/converter"
	pb "github.com/escoutdoor/kotopes/common/api/user/v1"
)

func (i *Implementation) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	usr, err := i.userService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{
		User: converter.ToPbFromUser(usr),
	}, nil
}
