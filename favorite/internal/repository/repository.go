package repository

import (
	"context"

	"github.com/escoutdoor/kotopes/favorite/internal/model"
)

type FavoriteRepository interface {
	Create(ctx context.Context, in *model.CreateFavorite) (string, error)
	IsPetFavoriteExists(ctx context.Context, petID, userID string) (bool, error)
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*model.Favorite, error)
	List(ctx context.Context, userID string) ([]*model.Favorite, error)
}
