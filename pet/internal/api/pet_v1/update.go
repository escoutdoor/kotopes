package pet_v1

import (
	"context"

	pb "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/pet/internal/converter"
	"github.com/escoutdoor/kotopes/pet/internal/model"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Update(ctx context.Context, req *pb.UpdateRequest) (*emptypb.Empty, error) {
	hasFields := i.hasFieldsForUpdating(req)
	if !hasFields {
		return nil, model.ErrNoFieldsForUpdating
	}

	err := i.petService.Update(ctx, converter.ToUpdatePetFromPb(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}

func (i *Implementation) hasFieldsForUpdating(req *pb.UpdateRequest) bool {
	return req.GetName() != nil || req.GetDescription() != nil || req.GetAge() != nil
}
