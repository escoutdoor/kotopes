package access

import (
	"context"

	"github.com/escoutdoor/kotopes/auth/internal/model"
	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/open-policy-agent/opa/rego"
)

const (
	userIDKey   = "user_id"
	methodKey   = "method"
	roleKey     = "role"
	endpointKey = "endpoint"
)

func (svc *service) CheckIsAllowed(ctx context.Context, in *model.AccessCheck) (*model.AccessInfo, error) {
	const op = "access_service.CheckIsAllowed"

	rg := rego.New(
		rego.Query("data.auth.allow"),
		rego.Compiler(svc.policy.Compiler),
		rego.Input(map[string]string{
			endpointKey: in.Endpoint,
			methodKey:   in.Method,
			userIDKey:   in.UserID,
			roleKey:     in.Role,
		}),
	)

	rs, err := rg.Eval(ctx)
	if err != nil || len(rs) == 0 {
		return nil, errwrap.Wrap(op, err)
	}

	if !rs.Allowed() {
		return &model.AccessInfo{IsAllowed: false}, nil
	}

	return &model.AccessInfo{IsAllowed: true}, nil
}
