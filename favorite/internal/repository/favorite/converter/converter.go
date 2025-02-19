package converter

import (
	"github.com/escoutdoor/kotopes/favorite/internal/model"
	repomodel "github.com/escoutdoor/kotopes/favorite/internal/repository/favorite/model"
)

func ToFavoriteFromRepo(repoFav *repomodel.Favorite) *model.Favorite {
	return &model.Favorite{
		ID:        repoFav.ID,
		UserID:    repoFav.UserID,
		PetID:     repoFav.PetID,
		CreatedAt: repoFav.CreatedAt,
	}
}

func ToFavoritesFromRepo(repoFavs []*repomodel.Favorite) []*model.Favorite {
	var favs []*model.Favorite
	for _, v := range repoFavs {
		favs = append(favs, ToFavoriteFromRepo(v))
	}

	return favs
}
