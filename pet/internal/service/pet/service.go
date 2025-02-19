package pet

import (
	"time"

	"github.com/escoutdoor/kotopes/common/pkg/db"
	"github.com/escoutdoor/kotopes/pet/internal/repository"
)

type service struct {
	txManager db.TxManager
	petRepo   repository.PetRepository
	petCache  repository.PetCache
	ttl       time.Duration
}

func NewService(
	petRepo repository.PetRepository,
	petCache repository.PetCache,
	txManager db.TxManager,
	ttl time.Duration,
) *service {
	return &service{
		petRepo:   petRepo,
		petCache:  petCache,
		txManager: txManager,
		ttl:       ttl,
	}
}
