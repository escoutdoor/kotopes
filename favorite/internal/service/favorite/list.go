package favorite

import (
	"context"

	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/favorite/internal/model"
)

func (svc *service) ListFavorites(ctx context.Context, in *model.ListFavorites) ([]*model.FavoritePet, error) {
	const op = "favorite_service.ListFavorites"

	favs, err := svc.favoriteRepo.List(ctx, in.UserID)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	var petIds []string
	for _, v := range favs {
		petIds = append(petIds, v.PetID)
	}

	var favorites []*model.FavoritePet
	pets, err := svc.petClient.ListPets(ctx, petIds)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	petM := map[string]*model.Pet{}
	for _, p := range pets {
		petM[p.ID] = p
	}

	for _, f := range favs {
		p := petM[f.PetID]
		fvpet := &model.FavoritePet{
			ID:        f.ID,
			UserID:    f.UserID,
			Pet:       p,
			CreatedAt: f.CreatedAt,
		}

		favorites = append(favorites, fvpet)
	}

	return favorites, nil
}
