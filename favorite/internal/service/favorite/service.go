package favorite

import (
	"github.com/escoutdoor/kotopes/common/pkg/db"
	grpc_client "github.com/escoutdoor/kotopes/favorite/internal/client/grpc"
	kafka_client "github.com/escoutdoor/kotopes/favorite/internal/client/kafka"
	"github.com/escoutdoor/kotopes/favorite/internal/repository"
)

type service struct {
	txManager          db.TxManager
	favoriteRepo       repository.FavoriteRepository
	petClient          grpc_client.PetClient
	notificationClient kafka_client.NotificationClient
}

func New(
	favoriteRepo repository.FavoriteRepository,
	txManager db.TxManager,
	petClient grpc_client.PetClient,
	notificationClient kafka_client.NotificationClient,
) *service {
	return &service{
		favoriteRepo:       favoriteRepo,
		txManager:          txManager,
		petClient:          petClient,
		notificationClient: notificationClient,
	}
}
