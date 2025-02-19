package auth_v1

import (
	"github.com/escoutdoor/kotopes/auth/internal/service"
	pb "github.com/escoutdoor/kotopes/common/api/auth/v1"
)

type Implementation struct {
	pb.UnimplementedAuthV1Server

	authService service.AuthService
}

func New(authService service.AuthService) *Implementation {
	return &Implementation{
		authService: authService,
	}
}
