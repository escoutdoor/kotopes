package auth

import (
	"github.com/escoutdoor/kotopes/auth/internal/repository"
	"github.com/escoutdoor/kotopes/auth/internal/utils/hasher"
	"github.com/escoutdoor/kotopes/auth/internal/utils/token"
)

type service struct {
	userRepo             repository.UserRepository
	accessTokenProvider  token.Provider
	refreshTokenProvider token.Provider
	pwHasher             hasher.Hasher
}

func New(
	userRepo repository.UserRepository,
	accessTokenProvider token.Provider,
	refreshTokenProvider token.Provider,
	pwHasher hasher.Hasher,
) *service {
	return &service{
		userRepo:             userRepo,
		accessTokenProvider:  accessTokenProvider,
		refreshTokenProvider: refreshTokenProvider,
		pwHasher:             pwHasher,
	}
}
