package favorite_v1

import (
	pb "github.com/escoutdoor/kotopes/common/api/favorite/v1"
	"github.com/escoutdoor/kotopes/favorite/internal/service"
)

type Implementation struct {
	pb.UnimplementedFavoriteV1Server

	favoriteService service.FavoriteService
}

func New(favoriteService service.FavoriteService) *Implementation {
	return &Implementation{
		favoriteService: favoriteService,
	}
}
