package user_v1

import (
	"github.com/escoutdoor/kotopes/auth/internal/service"
	pb "github.com/escoutdoor/kotopes/common/api/user/v1"
)

type Implementation struct {
	pb.UnimplementedUserV1Server

	userService service.UserService
}

func New(userService service.UserService) *Implementation {
	return &Implementation{
		userService: userService,
	}
}
