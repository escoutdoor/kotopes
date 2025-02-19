package favorite

import (
	"context"

	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/favorite/internal/model"
)

func (svc *service) Create(ctx context.Context, in *model.CreateFavorite) (string, error) {
	const op = "favorite_service.Create"

	pet, err := svc.petClient.Get(ctx, in.PetID)
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}

	isInFavorite, err := svc.favoriteRepo.IsPetFavoriteExists(ctx, in.PetID, in.UserID)
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}
	if isInFavorite {
		return "", errwrap.Wrap(op, model.ErrAlreadyInList)
	}

	var id string
	err = svc.txManager.ReadCommitted(ctx, func(ctx context.Context) error {
		id, err = svc.favoriteRepo.Create(ctx, in)
		if err != nil {
			return err
		}

		if pet.OwnerID != in.UserID {
			msg := &model.Notification{
				OwnerID: pet.OwnerID,
				PetID:   in.PetID,
				UserID:  in.UserID,
			}

			// mb not necessary??
			err = svc.notificationClient.Send(ctx, msg)
			if err != nil {
				return err
			}
		}

		return nil
	})
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}

	return id, nil
}
