package converter

import (
	"github.com/escoutdoor/kotopes/api-gateway/internal/controller/http/user/v1/dto"
	"github.com/escoutdoor/kotopes/api-gateway/internal/model"
)

func ToUpdateUserFromDTO(in *dto.UpdateUserRequest, userID string) *model.UpdateUser {
	return &model.UpdateUser{
		ID:        userID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  in.Password,
		Phone:     in.Phone,
		City:      in.City,
		Country:   in.Country,
	}
}

func ToCreateUserFromDTO(in *dto.CreateUserRequest) *model.CreateUser {
	return &model.CreateUser{
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Password:  in.Password,
	}
}

func ToDTOFromUsers(in []*model.User) []*dto.User {
	var users []*dto.User
	for _, u := range in {
		users = append(users, ToDTOFromUser(u))
	}

	return users
}

func ToDTOFromUser(in *model.User) *dto.User {
	return &dto.User{
		ID:        in.ID,
		FirstName: in.FirstName,
		LastName:  in.LastName,
		Email:     in.Email,
		Phone:     in.Phone,
		City:      in.City,
		Country:   in.Country,
		CreatedAt: in.CreatedAt,
	}
}
