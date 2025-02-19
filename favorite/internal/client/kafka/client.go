package kafka

import (
	"context"

	"github.com/escoutdoor/kotopes/favorite/internal/model"
)

type NotificationClient interface {
	Send(ctx context.Context, msg *model.Notification) error
}
