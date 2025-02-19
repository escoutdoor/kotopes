package converter

import (
	"github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/pet/v1/dto"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
)

func ToUpdatePetFromDTO(req *dto.UpdatePetRequest, petID, ownerID string) *model.UpdatePet {
	return &model.UpdatePet{
		ID:          petID,
		OwnerID:     ownerID,
		Name:        req.Name,
		Description: req.Description,
		Age:         req.Age,
	}
}

func ToCreatePetFromDTO(req *dto.CreatePetRequest, ownerID string) *model.CreatePet {
	return &model.CreatePet{
		Name:        req.Name,
		Description: req.Description,
		Age:         req.Age,
		OwnerID:     ownerID,
	}
}

func ToDTOFromPets(in []*model.Pet) []*dto.Pet {
	var pets []*dto.Pet
	for _, p := range in {
		pets = append(pets, ToDTOFromPet(p))
	}

	return pets
}

func ToDTOFromPet(in *model.Pet) *dto.Pet {
	return &dto.Pet{
		ID:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Age:         in.Age,
		OwnerID:     in.OwnerID,
		CreatedAt:   in.CreatedAt,
	}
}
