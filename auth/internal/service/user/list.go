package user

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/model"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
)

func (svc *service) List(ctx context.Context, userIDs []string) ([]*model.User, error) {
	const op = "user_service.List"

	users, err := svc.userRepo.List(ctx, userIDs)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return users, nil
}
