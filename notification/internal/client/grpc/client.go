package grpc

import (
	"context"

	"github.com/escoutdoor/kotopes/notification/internal/model"
)

type UserClient interface {
	List(ctx context.Context, userIDs []string) ([]*model.User, error)
	GetByID(ctx context.Context, id string) (*model.User, error)
}
