package converter

import (
	favoritepb "github.com/escoutdoor/kotopes/common/api/favorite/v1"
	petpb "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/favorite/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCreateFavoriteFromPb(req *favoritepb.CreateRequest) *model.CreateFavorite {
	return &model.CreateFavorite{
		PetID:  req.PetId,
		UserID: req.UserId,
	}
}

func ToPbFromFavorite(in *model.Favorite) *favoritepb.Favorite {
	return &favoritepb.Favorite{
		Id:        in.ID,
		UserId:    in.UserID,
		PetId:     in.PetID,
		CreatedAt: timestamppb.New(in.CreatedAt),
	}
}

func ToPbFromPet(in *model.Pet) *petpb.Pet {
	if in == nil {
		return nil
	}

	return &petpb.Pet{
		Id:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Age:         in.Age,
		OwnerId:     in.OwnerID,
		CreatedAt:   timestamppb.New(in.CreatedAt),
	}
}

func ToListFavoritesFromPb(req *favoritepb.ListFavoritesRequest) *model.ListFavorites {
	return &model.ListFavorites{
		UserID: req.UserId,
		Limit:  req.Limit,
		Offset: req.Offset,
	}
}

func ToDeleteFavoriteFromPb(req *favoritepb.DeleteRequest) *model.DeleteFavorite {
	return &model.DeleteFavorite{
		FavoriteID: req.FavoriteId,
		UserID:     req.UserId,
	}
}

func ToPbFromFavoritePet(in *model.FavoritePet) *favoritepb.ListFavoritesResponse_FavoritePet {
	return &favoritepb.ListFavoritesResponse_FavoritePet{
		Id:        in.ID,
		UserId:    in.UserID,
		Pet:       ToPbFromPet(in.Pet),
		CreatedAt: timestamppb.New(in.CreatedAt),
	}
}

func ToPbFromFavoritePets(in []*model.FavoritePet) []*favoritepb.ListFavoritesResponse_FavoritePet {
	var out []*favoritepb.ListFavoritesResponse_FavoritePet
	for _, fp := range in {
		out = append(out, ToPbFromFavoritePet(fp))
	}

	return out
}
