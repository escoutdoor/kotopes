package converter

import (
	"github.com/escoutdoor/kotopes/auth/internal/model"
	pb "github.com/escoutdoor/kotopes/common/api/auth/v1"
)

func ToCreateUserFromPb(req *pb.RegisterRequest) *model.CreateUser {
	return &model.CreateUser{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
	}
}

func ToLoginFromPb(req *pb.LoginRequest) *model.Login {
	return &model.Login{
		Email:    req.Email,
		Password: req.Password,
	}
}
