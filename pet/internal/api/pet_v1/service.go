package pet_v1

import (
	pb "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/pet/internal/service"
)

type Implementation struct {
	pb.UnimplementedPetV1Server

	petService service.PetService
}

func New(petService service.PetService) *Implementation {
	return &Implementation{
		petService: petService,
	}
}
