package converter

import (
	"github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/favorite/v1/dto"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
)

func ToDTOFromFavoritePets(in []*model.FavoritePet) []*dto.FavoritePet {
	var pets []*dto.FavoritePet
	for _, p := range in {
		pets = append(pets, ToDTOFromFavoritePet(p))
	}
	return pets
}

func ToDTOFromFavoritePet(in *model.FavoritePet) *dto.FavoritePet {
	return &dto.FavoritePet{
		ID:        in.ID,
		UserID:    in.UserID,
		Pet:       ToFavoriteDTOFromPet(in.Pet),
		CreatedAt: in.CreatedAt,
	}
}

func ToFavoriteDTOFromPet(in *model.Pet) *dto.Pet {
	return &dto.Pet{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Age:         in.Age,
		OwnerID:     in.OwnerID,
		CreatedAt:   in.CreatedAt,
	}
}
