package pet

import (
	"context"
	"errors"

	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/common/pkg/logger"
	rediscl "github.com/escoutdoor/kotopes/common/pkg/redis"
	"github.com/escoutdoor/kotopes/pet/internal/model"
	"go.uber.org/zap"
)

func (svc *service) Get(ctx context.Context, id string) (*model.Pet, error) {
	const op = "pet_service.Get"

	pet, err := svc.petCache.GetByID(ctx, id)
	if err == nil {
		return pet, nil
	} else {
		if !errors.Is(err, rediscl.ErrNotFound) {
			logger.Warn(
				ctx,
				"could not get pet from cache",
				zap.String("error", err.Error()),
			)
		}
	}

	pet, err = svc.petRepo.GetByID(ctx, id)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}

	err = svc.petCache.Set(ctx, pet, svc.ttl)
	if err != nil {
		logger.Warn(
			ctx,
			"could not set pet in cache",
			zap.String("error", err.Error()),
		)
	}
	return pet, nil
}
