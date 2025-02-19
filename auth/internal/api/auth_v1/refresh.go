package auth_v1

import (
	"context"

	pb "github.com/escoutdoor/kotopes/common/api/auth/v1"
)

func (i *Implementation) Refresh(ctx context.Context, req *pb.RefreshRequest) (*pb.RefreshResponse, error) {
	tokens, err := i.authService.Refresh(ctx, req.GetRefreshToken())
	if err != nil {
		return nil, err
	}

	return &pb.RefreshResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
