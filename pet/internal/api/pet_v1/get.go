package pet_v1

import (
	"context"

	pb "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/pet/internal/converter"
)

func (i *Implementation) Get(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	pet, err := i.petService.Get(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	return &pb.GetResponse{
		Pet: converter.ToPbFromPet(pet),
	}, nil
}
