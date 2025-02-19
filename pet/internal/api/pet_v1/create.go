package pet_v1

import (
	"context"

	pb "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/pet/internal/converter"
)

func (i *Implementation) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	id, err := i.petService.Create(ctx, converter.ToCreatePetFromPb(req))
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		Id: id,
	}, nil
}
