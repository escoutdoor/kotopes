package grpc

import (
	"context"

	"github.com/escoutdoor/kotopes/favorite/internal/model"
)

type PetClient interface {
	ListPets(ctx context.Context, ids []string) ([]*model.Pet, error)
	Get(ctx context.Context, id string) (*model.Pet, error)
}
