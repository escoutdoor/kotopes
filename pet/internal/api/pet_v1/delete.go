package pet_v1

import (
	"context"

	pb "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/pet/internal/converter"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	err := i.petService.Delete(ctx, converter.ToDeletePetFromPb(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
