package converter

import (
	"github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/auth/v1/dto"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
)

func ToLoginFromDTO(req *dto.LoginRequest) *model.Login {
	return &model.Login{
		Email:    req.Email,
		Password: req.Password,
	}
}

func ToRegisterFromDTO(req *dto.RegisterRequest) *model.Register {
	return &model.Register{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Password:  req.Password,
		Phone:     req.Phone,
	}
}
