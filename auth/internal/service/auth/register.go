package auth

import (
	"context"
	"errors"

	"github.com/escoutdoor/kotopes/auth/internal/model"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
)

func (svc *service) Register(ctx context.Context, in *model.CreateUser) (string, error) {
	const op = "auth_service.Register"

	user, err := svc.userRepo.GetByEmail(ctx, in.Email)
	if err != nil && !errors.Is(err, model.ErrUserNotFound) {
		return "", errwrap.Wrap(op, err)
	}
	if user != nil {
		return "", errwrap.Wrap(op, model.ErrUserAlreadyExists)
	}

	in.Password, err = svc.pwHasher.Hash(in.Password)
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}

	id, err := svc.userRepo.Create(ctx, in)
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}

	return id, nil
}
