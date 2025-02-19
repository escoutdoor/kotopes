package access

import (
	"context"

	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	accessapi "github.com/escoutdoor/kotopes/common/api/access/v1"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
)

type client struct {
	accessGRPCClient accessapi.AccessV1Client
}

func New(accessGRPCClient accessapi.AccessV1Client) *client {
	return &client{
		accessGRPCClient: accessGRPCClient,
	}
}

func (c *client) CheckIsAllowed(ctx context.Context, in *model.AccessCheck) (bool, error) {
	const op = "access_client.CheckIsAllowed"

	resp, err := c.accessGRPCClient.CheckIsAllowed(ctx, &accessapi.CheckRequest{
		Endpoint: in.Endpoint,
		Method:   in.Method,
		UserId:   in.UserID,
		Role:     in.Role,
	})
	if err != nil {
		return false, errwrap.Wrap(op, err)
	}

	return resp.IsAllowed, nil
}
