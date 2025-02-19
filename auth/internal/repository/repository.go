package repository

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/model"
)

type UserRepository interface {
	Create(ctx context.Context, in *model.CreateUser) (string, error)
	List(ctx context.Context, userIDs []string) ([]*model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	Update(ctx context.Context, in *model.UpdateUser) error
	Delete(ctx context.Context, id string) error
}
