package converter

import (
	petpb "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/favorite/internal/model"
)

func ToPetFromPb(in *petpb.Pet) *model.Pet {
	return &model.Pet{
		ID:          in.Id,
		Name:        in.Name,
		Description: in.Description,
		Age:         in.Age,
		OwnerID:     in.OwnerId,
		CreatedAt:   in.CreatedAt.AsTime(),
	}
}

func ToPetsFromPb(in []*petpb.Pet) []*model.Pet {
	var pets []*model.Pet
	for _, v := range in {
		pets = append(pets, ToPetFromPb(v))
	}

	return pets
}
