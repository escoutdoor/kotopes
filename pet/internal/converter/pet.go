package converter

import (
	pb "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/pet/internal/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToCreatePetFromPb(req *pb.CreateRequest) *model.CreatePet {
	return &model.CreatePet{
		Name:        req.Name,
		Description: req.Description,
		Age:         req.Age,
		OwnerID:     req.OwnerId,
	}
}

func ToPbFromPet(in *model.Pet) *pb.Pet {
	return &pb.Pet{
		Id:          in.ID,
		Name:        in.Name,
		Description: in.Description,
		Age:         in.Age,
		OwnerId:     in.OwnerID,
		CreatedAt:   timestamppb.New(in.CreatedAt),
	}
}

func ToPbFromPets(in []*model.Pet) []*pb.Pet {
	var pets []*pb.Pet
	for _, p := range in {
		pets = append(pets, ToPbFromPet(p))
	}

	return pets
}

func ToListPetsFromPb(req *pb.ListPetsRequest) *model.ListPets {
	return &model.ListPets{
		Limit:  req.Limit,
		Offset: req.Offset,
		PetIDs: req.PetIds,
	}
}

func ToUpdatePetFromPb(req *pb.UpdateRequest) *model.UpdatePet {
	out := &model.UpdatePet{
		ID:      req.Id,
		OwnerID: req.OwnerId,
	}

	if req.GetName() != nil {
		name := req.GetName().GetValue()
		out.Name = &name
	}

	if req.GetDescription() != nil {
		description := req.GetDescription().GetValue()
		out.Description = &description
	}

	if req.GetAge() != nil {
		age := req.GetAge().GetValue()
		out.Age = &age
	}
	return out
}

func ToDeletePetFromPb(req *pb.DeleteRequest) *model.DeletePet {
	return &model.DeletePet{
		ID:      req.Id,
		OwnerID: req.OwnerId,
	}
}
