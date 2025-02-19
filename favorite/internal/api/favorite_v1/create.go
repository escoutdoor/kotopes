package favorite_v1

import (
	"context"

	pb "github.com/escoutdoor/kotopes/common/api/favorite/v1"
	"github.com/escoutdoor/kotopes/favorite/internal/converter"
)

func (i *Implementation) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	id, err := i.favoriteService.Create(ctx, converter.ToCreateFavoriteFromPb(req))
	if err != nil {
		return nil, err
	}

	return &pb.CreateResponse{
		Id: id,
	}, nil
}
