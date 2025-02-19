package auth

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/model"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
)

func (svc *service) Validate(ctx context.Context, accessToken string) (*model.UserClaims, error) {
	const op = "auth_service.Refresh"

	claims, err := svc.accessTokenProvider.Verify(accessToken)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	return claims, nil
}
