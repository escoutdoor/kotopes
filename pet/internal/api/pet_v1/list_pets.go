package pet_v1

import (
	"context"

	pb "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/pet/internal/converter"
)

func (i *Implementation) ListPets(ctx context.Context, req *pb.ListPetsRequest) (*pb.ListPetsResponse, error) {
	pets, err := i.petService.ListPets(ctx, converter.ToListPetsFromPb(req))
	if err != nil {
		return nil, err
	}

	return &pb.ListPetsResponse{
		Pets: converter.ToPbFromPets(pets),
	}, nil
}
