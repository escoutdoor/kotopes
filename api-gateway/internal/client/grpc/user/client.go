package user

import (
	"context"

	"github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc/user/converter"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	userapi "github.com/escoutdoor/kotopes/common/api/user/v1"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
)

type client struct {
	userGRPCClient userapi.UserV1Client
}

func New(userGRPCClient userapi.UserV1Client) *client {
	return &client{
		userGRPCClient: userGRPCClient,
	}
}

func (c *client) Get(ctx context.Context, id string) (*model.User, error) {
	const op = "user_client.Get"

	resp, err := c.userGRPCClient.Get(ctx, &userapi.GetRequest{Id: id})
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return converter.ToUserFromApiResponse(resp.User), nil
}

func (c *client) List(ctx context.Context, in *model.ListUsers) ([]*model.User, error) {
	const op = "user_client.List"

	resp, err := c.userGRPCClient.List(ctx, &userapi.ListRequest{UserIds: in.UserIDs})
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	logger.Info(ctx, resp.Users)

	return converter.ToUsersFromApiResponse(resp.Users), nil
}

func (c *client) Update(ctx context.Context, in *model.UpdateUser) error {
	const op = "user_client.Update"

	_, err := c.userGRPCClient.Update(ctx, converter.ToUpdateUserRequestFromModel(in))
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}

func (c *client) Delete(ctx context.Context, id string) error {
	const op = "user_client.Delete"

	_, err := c.userGRPCClient.Delete(ctx, &userapi.DeleteRequest{Id: id})
	if err != nil {
		return errwrap.Wrap(op, err)
	}

	return nil
}
