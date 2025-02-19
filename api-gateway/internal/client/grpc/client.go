package grpc

import (
	"context"

	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
)

type AuthClient interface {
	Login(ctx context.Context, in *model.Login) (*model.AuthTokens, error)
	Register(ctx context.Context, in *model.Register) (string, error)
	Refresh(ctx context.Context, refreshToken string) (*model.AuthTokens, error)
	Validate(ctx context.Context, accessToken string) (*model.Token, error)
}

type AccessClient interface {
	CheckIsAllowed(ctx context.Context, in *model.AccessCheck) (bool, error)
}

type UserClient interface {
	Delete(ctx context.Context, id string) error
	Get(ctx context.Context, id string) (*model.User, error)
	List(ctx context.Context, in *model.ListUsers) ([]*model.User, error)
	Update(ctx context.Context, in *model.UpdateUser) error
}

type PetClient interface {
	Create(ctx context.Context, in *model.CreatePet) (string, error)
	Delete(ctx context.Context, petID string, ownerID string) error
	Get(ctx context.Context, id string) (*model.Pet, error)
	ListPets(ctx context.Context, in *model.ListPets) ([]*model.Pet, error)
	Update(ctx context.Context, in *model.UpdatePet) error
}

type FavoriteClient interface {
	Create(ctx context.Context, in *model.CreateFavorite) (string, error)
	Delete(ctx context.Context, in *model.DeleteFavorite) error
	ListFavorites(ctx context.Context, in *model.ListFavorites) (*model.FavoriteList, error)
}
