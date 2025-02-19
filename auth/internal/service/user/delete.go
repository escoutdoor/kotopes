package user

import (
	"context"

	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
)

func (svc *service) Delete(ctx context.Context, id string) error {
	const op = "user_service.Delete"

	err := svc.userRepo.Delete(ctx, id)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}
