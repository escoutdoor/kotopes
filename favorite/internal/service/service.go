package service

import (
	"context"

	"github.com/escoutdoor/kotopes/favorite/internal/model"
)

type FavoriteService interface {
	Create(ctx context.Context, in *model.CreateFavorite) (string, error)
	Delete(ctx context.Context, in *model.DeleteFavorite) error
	ListFavorites(ctx context.Context, in *model.ListFavorites) ([]*model.FavoritePet, error)
}
