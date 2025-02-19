package pet

import (
	"context"

	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/pet/internal/model"
)

func (svc *service) ListPets(ctx context.Context, in *model.ListPets) ([]*model.Pet, error) {
	const op = "pet_service.ListPets"

	pets, err := svc.petRepo.ListPets(ctx, in)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}
	return pets, nil
}
