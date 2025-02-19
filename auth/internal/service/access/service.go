package access

import (
	"github.com/escoutdoor/kotopes/auth/internal/utils/policy"
)

type service struct {
	policy *policy.Policy
}

func New(policy *policy.Policy) *service {
	return &service{
		policy: policy,
	}
}
