package state

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type KeyValueStateStore interface {
	GetName() string
	Get(ctx context.Context, key string) (string, error)
	Put(ctx context.Context, key string, val string) error
	Exists(ctx context.Context, key string) (string, error)
}

type RedisKeyValueStateStore struct {
	KeyValueStateStore
	r    *redis.Client
	name string
}

func NewRedisKeyValueStateStore() *RedisKeyValueStateStore {
	return &RedisKeyValueStateStore{}
}

func (s *RedisKeyValueStateStore) Get(ctx context.Context, key string) (string, error) {
	return s.r.Get(ctx, key).Result()
}

func (s *RedisKeyValueStateStore) Put(ctx context.Context, key string, val any) error {
	return s.r.Set(ctx, key, val, 0).Err()
}

func (s *RedisKeyValueStateStore) Exists(ctx context.Context, key string) (bool, error) {
	count, err := s.r.Exists(ctx, key).Result()
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
