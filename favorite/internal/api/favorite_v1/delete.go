package favorite_v1

import (
	"context"

	pb "github.com/escoutdoor/kotopes/common/api/favorite/v1"
	"github.com/escoutdoor/kotopes/favorite/internal/converter"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (i *Implementation) Delete(ctx context.Context, req *pb.DeleteRequest) (*emptypb.Empty, error) {
	err := i.favoriteService.Delete(ctx, converter.ToDeleteFavoriteFromPb(req))
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}
