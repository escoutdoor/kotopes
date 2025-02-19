package user

import (
	"context"

	userpb "github.com/escoutdoor/kotopes/common/api/user/v1"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/notification/internal/client/grpc/user/converter"
	"github.com/escoutdoor/kotopes/notification/internal/model"
)

type client struct {
	userGRPCClient userpb.UserV1Client
}

func New(userGRPCClient userpb.UserV1Client) *client {
	return &client{
		userGRPCClient: userGRPCClient,
	}
}

func (cl *client) List(ctx context.Context, userIDs []string) ([]*model.User, error) {
	const op = "user_client.List"

	resp, err := cl.userGRPCClient.List(ctx, &userpb.ListRequest{UserIds: userIDs})
	if err != nil {
		return nil, err
	}

	return converter.ToUsersFromPb(resp.Users), nil
}

func (cl *client) GetByID(ctx context.Context, id string) (*model.User, error) {
	const op = "user_client.GetByID"

	resp, err := cl.userGRPCClient.Get(ctx, &userpb.GetRequest{Id: id})
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToUserFromPb(resp.User), nil
}
