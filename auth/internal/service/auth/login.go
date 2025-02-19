package auth

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/model"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
)

func (svc *service) Login(ctx context.Context, in *model.Login) (*model.Tokens, error) {
	const op = "auth_service.Login"

	user, err := svc.userRepo.GetByEmail(ctx, in.Email)
	if err != nil {
		return nil, errwrap.Wrap(op, model.ErrInvalidEmailOrPassword)
	}

	ok := svc.pwHasher.Compare(in.Password, user.Password)
	if !ok {
		return nil, errwrap.Wrap(op, model.ErrInvalidEmailOrPassword)
	}

	accessToken, err := svc.accessTokenProvider.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	refreshToken, err := svc.refreshTokenProvider.GenerateToken(user.ID, user.Role)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return &model.Tokens{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
