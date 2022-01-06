package storage

import (
	"context"
	"errors"
	"time"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

type Redis struct {
	client *redis.Client
}

func NewRedis(host string, password string) Redis {
	return Redis{
		redis.NewClient(&redis.Options{
			Addr:     host,
			Password: password,
			DB:       0,
		}),
	}
}

func (r Redis) Set(key string, value string, expiresAt time.Duration) error {
	err := r.client.Set(ctx, key, value, expiresAt).Err()
	if err != nil {
		return errors.New("error when trying to save")
	}

	return nil
}

func (r Redis) Get(key string) (string, error) {
	v := r.client.Get(ctx, key)

	return v.Result()
}
