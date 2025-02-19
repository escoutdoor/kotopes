package converter

import (
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	petapi "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToCreatePetRequestFromModel(in *model.CreatePet) *petapi.CreateRequest {
	return &petapi.CreateRequest{
		Name:        in.Name,
		Description: in.Description,
		Age:         in.Age,
		OwnerId:     in.OwnerID,
	}
}

func ToUpdatePetRequestFromModel(in *model.UpdatePet) *petapi.UpdateRequest {
	out := &petapi.UpdateRequest{
		Id:      in.ID,
		OwnerId: in.OwnerID,
	}

	if in.Name != nil {
		out.Name = wrapperspb.String(*in.Name)
	}
	if in.Description != nil {
		out.Description = wrapperspb.String(*in.Description)
	}
	if in.Age != nil {
		out.Age = wrapperspb.Int32(*in.Age)
	}

	return out
}

func ToListPetsRequestFromModel(in *model.ListPets) *petapi.ListPetsRequest {
	return &petapi.ListPetsRequest{
		Limit:  in.Limit,
		Offset: in.Offset,
		PetIds: in.PetIDs,
	}
}

func ToPetFromApiResponse(resp *petapi.Pet) *model.Pet {
	return &model.Pet{
		ID:          resp.Id,
		Name:        resp.Name,
		Description: resp.Description,
		Age:         resp.Age,
		OwnerID:     resp.OwnerId,
		CreatedAt:   resp.CreatedAt.AsTime(),
	}
}

func ToPetsFromApiResponse(resp []*petapi.Pet) []*model.Pet {
	var pets []*model.Pet
	for _, p := range resp {
		pets = append(pets, ToPetFromApiResponse(p))
	}

	return pets
}
