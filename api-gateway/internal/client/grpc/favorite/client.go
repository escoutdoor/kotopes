package favorite

import (
	"context"

	"github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc/favorite/converter"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	favoriteapi "github.com/escoutdoor/kotopes/common/api/favorite/v1"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
)

type client struct {
	favoriteGRPCClient favoriteapi.FavoriteV1Client
}

func New(favoriteGRPCClient favoriteapi.FavoriteV1Client) *client {
	return &client{
		favoriteGRPCClient: favoriteGRPCClient,
	}
}

func (c *client) Create(ctx context.Context, in *model.CreateFavorite) (string, error) {
	const op = "favorite_service.Create"

	resp, err := c.favoriteGRPCClient.Create(ctx, converter.ToCreateRequestFromModel(in))
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}

	return resp.Id, nil
}

func (c *client) Delete(ctx context.Context, in *model.DeleteFavorite) error {
	const op = "favorite_service.Delete"

	_, err := c.favoriteGRPCClient.Delete(ctx, converter.ToDeleteRequestFromModel(in))
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}

func (c *client) ListFavorites(ctx context.Context, in *model.ListFavorites) (*model.FavoriteList, error) {
	const op = "favorite_service.ListFavorites"

	resp, err := c.favoriteGRPCClient.ListFavorites(ctx, converter.ToListFavoritesRequestFromModel(in))
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	out := &model.FavoriteList{
		Favorites: converter.ToFavoritePetsFromApiResponse(resp.Favorites),
		Total:     resp.Total,
	}
	return out, nil
}
