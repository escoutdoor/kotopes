package auth

import (
	"context"

	"github.com/escoutdoor/kotopes/api-gateway/internal/client/grpc/auth/converter"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	authapi "github.com/escoutdoor/kotopes/common/api/auth/v1"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
)

type client struct {
	authGRPCClient authapi.AuthV1Client
}

func New(authGRPCClient authapi.AuthV1Client) *client {
	return &client{
		authGRPCClient: authGRPCClient,
	}
}

func (c *client) Login(ctx context.Context, in *model.Login) (*model.AuthTokens, error) {
	const op = "auth_client.Login"

	resp, err := c.authGRPCClient.Login(ctx, converter.ToLoginRequestFromModel(in))
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	out := &model.AuthTokens{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}
	return out, nil
}

func (c *client) Register(ctx context.Context, in *model.Register) (string, error) {
	const op = "auth_client.Register"

	resp, err := c.authGRPCClient.Register(ctx, converter.ToRegisterRequestFromModel(in))
	if err != nil {
		return "", errwrap.Wrap(op, err)
	}

	return resp.Id, nil
}

func (c *client) Refresh(ctx context.Context, refreshToken string) (*model.AuthTokens, error) {
	const op = "auth_client.Refresh"

	resp, err := c.authGRPCClient.Refresh(ctx, &authapi.RefreshRequest{RefreshToken: refreshToken})
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	out := &model.AuthTokens{
		AccessToken:  resp.AccessToken,
		RefreshToken: resp.RefreshToken,
	}
	return out, nil
}

func (c *client) Validate(ctx context.Context, accessToken string) (*model.Token, error) {
	const op = "auth_client.Validate"

	resp, err := c.authGRPCClient.Validate(ctx, &authapi.ValidateRequest{AccessToken: accessToken})
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return &model.Token{
		UserID: resp.Id,
		Role:   resp.Role,
	}, nil
}
