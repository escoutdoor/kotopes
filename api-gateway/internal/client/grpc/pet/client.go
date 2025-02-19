package pet

import (
	"context"

	"github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc/pet/converter"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	petapi "github.com/escoutdoor/kotopes/common/api/pet/v1"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
)

type client struct {
	petGRPCClient petapi.PetV1Client
}

func New(petGRPCClient petapi.PetV1Client) *client {
	return &client{
		petGRPCClient: petGRPCClient,
	}
}

func (c *client) Create(ctx context.Context, in *model.CreatePet) (string, error) {
	const op = "pet_client.Create"

	resp, err := c.petGRPCClient.Create(ctx, converter.ToCreatePetRequestFromModel(in))
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}

	return resp.Id, nil
}

func (c *client) Get(ctx context.Context, id string) (*model.Pet, error) {
	const op = "pet_client.Get"

	resp, err := c.petGRPCClient.Get(ctx, &petapi.GetRequest{Id: id})
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToPetFromApiResponse(resp.Pet), nil
}

func (c *client) ListPets(ctx context.Context, in *model.ListPets) ([]*model.Pet, error) {
	const op = "pet_client.ListPets"

	resp, err := c.petGRPCClient.ListPets(ctx, converter.ToListPetsRequestFromModel(in))
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToPetsFromApiResponse(resp.Pets), nil
}

func (c *client) Update(ctx context.Context, in *model.UpdatePet) error {
	const op = "pet_client.Update"

	_, err := c.petGRPCClient.Update(ctx, converter.ToUpdatePetRequestFromModel(in))
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}

func (c *client) Delete(ctx context.Context, petID, ownerID string) error {
	const op = "pet_client.Delete"

	_, err := c.petGRPCClient.Delete(ctx,
		&petapi.DeleteRequest{Id: petID, OwnerId: ownerID},
	)
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}
