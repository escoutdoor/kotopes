package converter

import (
	pet_converter "github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc/pet/converter"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	favoriteapi "github.com/escoutdoor/kotopes/common/api/favorite/v1"
)

func ToFavoritePetFromApiResponse(resp *favoriteapi.ListFavoritesResponse_FavoritePet) *model.FavoritePet {
	return &model.FavoritePet{
		ID:        resp.Id,
		UserID:    resp.UserId,
		Pet:       pet_converter.ToPetFromApiResponse(resp.Pet),
		CreatedAt: resp.CreatedAt.AsTime(),
	}
}

func ToFavoritePetsFromApiResponse(resp []*favoriteapi.ListFavoritesResponse_FavoritePet) []*model.FavoritePet {
	var favoritePets []*model.FavoritePet
	for _, fp := range resp {
		favoritePets = append(favoritePets, ToFavoritePetFromApiResponse(fp))
	}

	return favoritePets
}

func ToCreateRequestFromModel(in *model.CreateFavorite) *favoriteapi.CreateRequest {
	return &favoriteapi.CreateRequest{
		PetId:  in.PetID,
		UserId: in.UserID,
	}
}

func ToDeleteRequestFromModel(in *model.DeleteFavorite) *favoriteapi.DeleteRequest {
	return &favoriteapi.DeleteRequest{
		FavoriteId: in.FavoriteID,
		UserId:     in.UserID,
	}
}

func ToListFavoritesRequestFromModel(in *model.ListFavorites) *favoriteapi.ListFavoritesRequest {
	return &favoriteapi.ListFavoritesRequest{
		UserId: in.UserID,
		Limit:  in.Limit,
		Offset: in.Offset,
	}
}
