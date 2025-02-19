package auth

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/model"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
)

func (svc *service) Refresh(ctx context.Context, refreshToken string) (*model.Tokens, error) {
	const op = "auth_service.Refresh"

	claims, err := svc.refreshTokenProvider.Verify(refreshToken)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	accessToken, err := svc.accessTokenProvider.GenerateToken(claims.ID, claims.Role)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	newRefreshToken, err := svc.refreshTokenProvider.GenerateToken(claims.ID, claims.Role)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return &model.Tokens{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}
