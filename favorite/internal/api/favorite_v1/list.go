package favorite_v1

import (
	"context"

	pb "github.com/escoutdoor/kotopes/common/api/favorite/v1"
	"github.com/escoutdoor/kotopes/favorite/internal/converter"
)

func (i *Implementation) ListFavorites(ctx context.Context, req *pb.ListFavoritesRequest) (*pb.ListFavoritesResponse, error) {
	favs, err := i.favoriteService.ListFavorites(ctx, converter.ToListFavoritesFromPb(req))
	if err != nil {
		return nil, err
	}

	return &pb.ListFavoritesResponse{
		Favorites: converter.ToPbFromFavoritePets(favs),
		Total:     int32(len(favs)),
	}, nil
}
