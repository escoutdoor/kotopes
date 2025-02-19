package pet

import (
	"context"

	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/pet/internal/model"
)

func (svc *service) Delete(ctx context.Context, in *model.DeletePet) error {
	const op = "pet_service.Delete"

	pet, err := svc.petRepo.GetByID(ctx, in.ID)
	if err != nil {
		return errwrap.Wrap(op, err)
	}
	if pet.OwnerID != in.OwnerID {
		return errwrap.Wrap(op, model.ErrNotPetOwner)
	}

	err = svc.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		err := svc.petRepo.Delete(ctx, in.ID)
		if err != nil {
			return err
		}

		err = svc.petCache.Delete(ctx, in.ID)
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}
