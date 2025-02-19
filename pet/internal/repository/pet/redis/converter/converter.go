package converter

import (
	"github.com/escoutdoor/kotopes/pet/internal/model"
	cachemodel "github.com/escoutdoor/kotopes/pet/internal/repository/pet/redis/model"
)

func ToPetFromRepo(cachedPet *cachemodel.Pet) *model.Pet {
	return &model.Pet{
		ID:          cachedPet.ID,
		Name:        cachedPet.Name,
		Description: cachedPet.Description,
		Age:         cachedPet.Age,
		OwnerID:     cachedPet.OwnerID,
		CreatedAt:   cachedPet.CreatedAt,
	}
}

func ToRepoFromPet(pet *model.Pet) *cachemodel.Pet {
	return &cachemodel.Pet{
		ID:          pet.ID,
		Name:        pet.Name,
		Description: pet.Description,
		Age:         pet.Age,
		OwnerID:     pet.OwnerID,
		CreatedAt:   pet.CreatedAt,
	}
}
