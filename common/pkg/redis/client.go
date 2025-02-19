package redis

import (
	"context"
	"errors"
	"time"

	"github.com/escoutdoor/kotopes/common/pkg/redis/config"
	"github.com/redis/go-redis/v9"
)

var (
	ErrNotFound = errors.New("not found")
)

type Client interface {
	Ping(ctx context.Context) error
	Close() error
	Delete(ctx context.Context, key string) error
	Get(ctx context.Context, key string, value interface{}) error
	Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error
}

type client struct {
	redisClient *redis.Client
}

func NewClient(ctx context.Context, cfg config.RedisCfg) (Client, error) {
	redisClient := redis.NewClient(&redis.Options{
		Addr:         cfg.Addr(),
		DB:           cfg.DB(),
		Password:     cfg.Password(),
		MaxIdleConns: cfg.MaxIdle(),
		DialTimeout:  cfg.ConnTimeout(),
	})

	return &client{
		redisClient: redisClient,
	}, nil
}

func (cl *client) Ping(ctx context.Context) error {
	return cl.redisClient.Ping(ctx).Err()
}

func (cl *client) Close() error {
	return cl.redisClient.Close()
}

func (cl *client) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	err := cl.redisClient.Set(ctx, key, value, expiration).Err()
	if err != nil {
		return err
	}

	return nil
}

func (cl *client) Get(ctx context.Context, key string, value interface{}) error {
	err := cl.redisClient.Get(ctx, key).Scan(value)
	if err != nil {
		if errors.Is(err, redis.Nil) {
			return ErrNotFound
		}
		return err
	}

	return nil
}

func (cl *client) Delete(ctx context.Context, key string) error {
	err := cl.redisClient.Del(ctx, key).Err()
	if err != nil {
		return err
	}

	return nil
}
