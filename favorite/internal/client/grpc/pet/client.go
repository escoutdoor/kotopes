package pet

import (
	"context"

	petpb "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	grpc_client "github.com/escoutdoor/kotopes/favorite/internal/client/grpc"
	"github.com/escoutdoor/kotopes/favorite/internal/client/grpc/pet/converter"
	"github.com/escoutdoor/kotopes/favorite/internal/model"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type petClient struct {
	petGRPCClient petpb.PetV1Client
}

func New(petGRPCClient petpb.PetV1Client) grpc_client.PetClient {
	return &petClient{
		petGRPCClient: petGRPCClient,
	}
}

func (cl *petClient) ListPets(ctx context.Context, ids []string) ([]*model.Pet, error) {
	const op = "pet_client.ListPets"

	resp, err := cl.petGRPCClient.ListPets(ctx,
		&petpb.ListPetsRequest{PetIds: ids})
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToPetsFromPb(resp.Pets), nil
}

func (cl *petClient) Get(ctx context.Context, id string) (*model.Pet, error) {
	const op = "pet_client.Get"

	resp, err := cl.petGRPCClient.Get(ctx, &petpb.GetRequest{Id: id})
	if err != nil {
		st, ok := status.FromError(err)
		if ok && st.Code() == codes.NotFound {
			return nil, errwrap.Wrap(op, model.ErrFavoriteNotFound)
		}

		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToPetFromPb(resp.Pet), nil
}
