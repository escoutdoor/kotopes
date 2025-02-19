package access_v1

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/converter"
	pb "github.com/escoutdoor/kotopes/common/api/access/v1"
)

func (i *Implementation) CheckIsAllowed(ctx context.Context, req *pb.CheckRequest) (*pb.CheckResponse, error) {
	accessInfo, err := i.accessService.CheckIsAllowed(ctx, converter.ToAccessCheckFromPb(req))
	if err != nil {
		return nil, err
	}

	return &pb.CheckResponse{IsAllowed: accessInfo.IsAllowed}, nil
}
