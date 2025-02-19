package converter

import (
	"github.com/escoutdoor/kotopes/auth/internal/model"
	pb "github.com/escoutdoor/kotopes/common/api/access/v1"
)

func ToAccessCheckFromPb(req *pb.CheckRequest) *model.AccessCheck {
	return &model.AccessCheck{
		Endpoint: req.Endpoint,
		Method:   req.Method,
		UserID:   req.UserId,
		Role:     req.Role,
	}
}
