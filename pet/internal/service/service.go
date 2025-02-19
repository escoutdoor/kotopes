package service

import (
	"context"

	"github.com/escoutdoor/kotopes/pet/internal/model"
)

type PetService interface {
	Create(ctx context.Context, in *model.CreatePet) (string, error)
	Get(ctx context.Context, id string) (*model.Pet, error)
	ListPets(ctx context.Context, in *model.ListPets) ([]*model.Pet, error)
	Update(ctx context.Context, in *model.UpdatePet) error
	Delete(ctx context.Context, in *model.DeletePet) error
}
