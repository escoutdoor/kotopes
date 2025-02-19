package converter

import (
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
	userapi "github.com/escoutdoor/kotopes/common/api/user/v1"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func ToUserFromApiResponse(resp *userapi.User) *model.User {
	return &model.User{
		ID:        resp.Id,
		FirstName: resp.FirstName,
		LastName:  resp.LastName,
		Email:     resp.Email,
		Phone:     resp.Phone,
		City:      resp.City,
		Country:   resp.Country,
		CreatedAt: resp.CreatedAt.AsTime(),
	}
}

func ToUsersFromApiResponse(resp []*userapi.User) []*model.User {
	var users []*model.User
	for _, u := range resp {
		users = append(users, ToUserFromApiResponse(u))
	}

	return users
}

func ToUpdateUserRequestFromModel(in *model.UpdateUser) *userapi.UpdateRequest {
	out := &userapi.UpdateRequest{
		Id: in.ID,
	}

	if in.FirstName != nil {
		out.FirstName = wrapperspb.String(*in.FirstName)
	}
	if in.LastName != nil {
		out.LastName = wrapperspb.String(*in.LastName)
	}
	if in.Email != nil {
		out.Email = wrapperspb.String(*in.Email)
	}
	if in.Password != nil {
		out.Password = wrapperspb.String(*in.Password)
	}
	if in.Phone != nil {
		out.Phone = wrapperspb.String(*in.Phone)
	}
	if in.City != nil {
		out.City = wrapperspb.String(*in.City)
	}
	if in.Country != nil {
		out.Country = wrapperspb.String(*in.Country)
	}

	return out
}
