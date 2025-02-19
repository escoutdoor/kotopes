package service

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/model"
)

type AuthService interface {
	Register(ctx context.Context, in *model.CreateUser) (string, error)
	Login(ctx context.Context, in *model.Login) (*model.Tokens, error)
	Refresh(ctx context.Context, refreshToken string) (*model.Tokens, error)
	Validate(ctx context.Context, accessToken string) (*model.UserClaims, error)
}

type UserService interface {
	List(ctx context.Context, userIDs []string) ([]*model.User, error)
	Get(ctx context.Context, id string) (*model.User, error)
	Update(ctx context.Context, in *model.UpdateUser) error
	Delete(ctx context.Context, id string) error
}

type AccessService interface {
	CheckIsAllowed(ctx context.Context, in *model.AccessCheck) (*model.AccessInfo, error)
}
