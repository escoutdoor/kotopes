package user

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/model"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
)

func (svc *service) Update(ctx context.Context, in *model.UpdateUser) error {
	const op = "user_service.Update"

	err := svc.userRepo.Update(ctx, in)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}
