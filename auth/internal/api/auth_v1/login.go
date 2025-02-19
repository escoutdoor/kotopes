package auth_v1

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/converter"
	pb "github.com/escoutdoor/kotopes/common/api/auth/v1"
)

func (i *Implementation) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	tokens, err := i.authService.Login(ctx, converter.ToLoginFromPb(req))
	if err != nil {
		return nil, err
	}

	return &pb.LoginResponse{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}, nil
}
