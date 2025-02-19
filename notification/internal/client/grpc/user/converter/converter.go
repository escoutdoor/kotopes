package converter

import (
	userpb "github.com/escoutdoor/kotopes/common/api/user/v1"
	"github.com/escoutdoor/kotopes/notification/internal/model"
)

func ToUserFromPb(in *userpb.User) *model.User {
	return &model.User{
		ID:        in.Id,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Phone:     in.Phone,
		City:      in.City,
		Country:   in.Country,
		CreatedAt: in.CreatedAt.AsTime(),
	}
}

func ToUsersFromPb(in []*userpb.User) []*model.User {
	var users []*model.User
	for _, u := range in {
		users = append(users, ToUserFromPb(u))
	}

	return users
}
