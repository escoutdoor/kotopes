package redis

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/escoutdoor/kotopes/common/pkg/errwrap"
	"github.com/escoutdoor/kotopes/common/pkg/redis"
	"github.com/escoutdoor/kotopes/pet/internal/model"
	"github.com/escoutdoor/kotopes/pet/internal/repository/pet/redis/converter"
	cachemodel "github.com/escoutdoor/kotopes/pet/internal/repository/pet/redis/model"
)

type repository struct {
	redisClient redis.Client
}

func New(redisClient redis.Client) *repository {
	return &repository{
		redisClient: redisClient,
	}
}

func (r *repository) Set(ctx context.Context, in *model.Pet, expiration time.Duration) error {
	const op = "pet_cache.Set"
	pet := converter.ToRepoFromPet(in)

	err := r.redisClient.Set(ctx, r.generateKey(in.ID), pet, expiration)
	if err != nil {
		return errwrap.Wrap(op, err)
	}
	return nil
}

func (r *repository) Delete(ctx context.Context, id string) error {
	const op = "pet_cache.Delete"

	err := r.redisClient.Delete(ctx, r.generateKey(id))
	if err != nil {
		return errwrap.Wrap(op, err)
	}
	return nil
}

func (r *repository) GetByID(ctx context.Context, id string) (*model.Pet, error) {
	const op = "pet_cache.GetByID"

	pet := cachemodel.Pet{}
	err := r.redisClient.Get(ctx, r.generateKey(id), &pet)
	if err != nil {
		return nil, errwrap.Wrap(op, err)
	}
	return converter.ToPetFromRepo(&pet), nil
}

func (r *repository) generateKey(ids ...string) string {
	return fmt.Sprintf("pet:%s", strings.Join(ids, ","))
}
