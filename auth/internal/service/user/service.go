package user

import "github.com/escoutdoor/kotopes/auth/internal/repository"

type service struct {
	userRepo repository.UserRepository
}

func New(userRepo repository.UserRepository) *service {
	return &service{
		userRepo: userRepo,
	}
}
