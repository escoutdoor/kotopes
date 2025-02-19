package pet

import (
	"context"

	errwrap "github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/pet/internal/model"
)

func (svc *service) Create(ctx context.Context, in *model.CreatePet) (string, error) {
	const op = "pet_service.Create"

	id, err := svc.petRepo.Create(ctx, in)
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}
	return id, nil
}
