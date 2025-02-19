package user

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/model"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
)

func (svc *service) Get(ctx context.Context, id string) (*model.User, error) {
	const op = "user_service.Get"

	user, err := svc.userRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return user, nil
}
