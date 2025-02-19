package converter

import (
	"database/sql"

	"github.com/escoutdoor/kotopes/auth/internal/model"
	repomodel "github.com/escoutdoor/kotopes/auth/internal/repository/user/model"
)

func ToUserFromRepo(repoUser *repomodel.User) *model.User {
	return &model.User{
		ID:        repoUser.ID,
		FirstName: repoUser.FirstName,
		LastName:  repoUser.LastName,
		Email:     repoUser.Email,
		Password:  repoUser.Password,
		Role:      repoUser.Role,
		Phone:     fromNullStrToStr(repoUser.Phone),
		City:      fromNullStrToStr(repoUser.City),
		Country:   fromNullStrToStr(repoUser.Country),
		CreatedAt: repoUser.CreatedAt,
	}
}

func ToUsersFromRepo(repoUsers []*repomodel.User) []*model.User {
	var users []*model.User
	for _, u := range repoUsers {
		users = append(users, ToUserFromRepo(u))
	}

	return users
}

func fromNullStrToStr(v sql.NullString) string {
	if v.Valid {
		return v.String
	}
	return ""
}
