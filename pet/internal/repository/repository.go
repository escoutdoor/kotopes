package repository

import (
	"context"
	"time"

	"github.com/escoutdoor/kotopes/pet/internal/model"
)

type PetRepository interface {
	Create(ctx context.Context, in *model.CreatePet) (string, error)
	GetByID(ctx context.Context, id string) (*model.Pet, error)
	ListPets(ctx context.Context, in *model.ListPets) ([]*model.Pet, error)
	Update(ctx context.Context, in *model.UpdatePet) error
	Delete(ctx context.Context, id string) error
}

type PetCache interface {
	Set(ctx context.Context, in *model.Pet, expiration time.Duration) error
	Delete(ctx context.Context, id string) error
	GetByID(ctx context.Context, id string) (*model.Pet, error)
}
