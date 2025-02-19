package favorite

import (
	"context"

	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/favorite/internal/model"
)

func (svc *service) Delete(ctx context.Context, in *model.DeleteFavorite) error {
	const op = "favorite_service.Delete"

	favorite, err := svc.favoriteRepo.GetByID(ctx, in.FavoriteID)
	if err != nil {
		return errwrap.Wrap(op, err)
	}
	if favorite.UserID != in.UserID {
		return errwrap.Wrap(op, model.ErrNotOwnerOfFavorite)
	}

	err = svc.favoriteRepo.Delete(ctx, in.FavoriteID)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}
