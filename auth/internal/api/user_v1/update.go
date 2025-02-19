package user_v1

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/converter"
	pb "github.com/escoutdoor/kotopes/common/api/user/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	err := i.userService.Update(ctx, converter.ToUpdateUserFromPb(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
