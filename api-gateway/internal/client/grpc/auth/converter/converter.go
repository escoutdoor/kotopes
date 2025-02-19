package converter

import (
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	authapi "github.com/escoutdoor/kotopes/common/api/auth/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToLoginRequestFromModel(in *model.Login) *authapi.LoginRequest {
	return &authapi.LoginRequest{
		Email:    in.Email,
		Password: in.Password,
	}
}

func ToRegisterRequestFromModel(in *model.Register) *authapi.RegisterRequest {
	out := &authapi.RegisterRequest{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  in.Password,
	}

	if in.Phone != nil {
		out.Phone = wrapperspb.String(*in.Phone)
	}
	return out
}
