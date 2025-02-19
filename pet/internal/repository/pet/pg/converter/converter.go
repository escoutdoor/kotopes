package converter

import (
	"github.com/escoutdoor/kotopes/pet/internal/model"
	repomodel "github.com/escoutdoor/kotopes/pet/internal/repository/pet/pg/model"
)

func ToPetFromRepo(repoPet *repomodel.Pet) *model.Pet {
	return &model.Pet{
		ID:          repoPet.ID,
		Name:        repoPet.Name,
		Description: repoPet.Description,
		Age:         repoPet.Age,
		OwnerID:     repoPet.OwnerID,
		CreatedAt:   repoPet.CreatedAt,
	}
}

func ToPetsFromRepo(repoPets []*repomodel.Pet) []*model.Pet {
	var pets []*model.Pet
	for _, v := range repoPets {
		pets = append(pets, ToPetFromRepo(v))
	}

	return pets
}
