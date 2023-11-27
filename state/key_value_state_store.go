package state

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type KeyValueStateStore interface {
	StateStore
	Get(ctx context.Context, key string) (string, error)
	Put(ctx context.Context, key string, val string) error
	Exists(ctx context.Context, key string) (string, error)
}

type RedisKeyValueStateStore struct {
	KeyValueStateStore

	name     string
	isGlobal bool
	address  string

	r *redis.Client
}

func newRedisKeyValueStateStore(name string, isGlobal bool, address string) (*RedisKeyValueStateStore, error) {
	// TODO: what happens if connection failed?
	cli := redis.NewClient(&redis.Options{
		Addr: address,
	})
	if err := cli.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}
	s := &RedisKeyValueStateStore{
		name:     name,
		isGlobal: isGlobal,
		address:  address,
		r:        cli,
	}
	return s, nil
}

func (s *RedisKeyValueStateStore) Name() string {
	return s.name
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

type RedisKeyValueStateStoreBuilder struct {
	Builder
	name     string
	isGlobal bool
	address  string
}

func NewRedisKeyValueStateStoreBuilder(name string, isGlobal bool, address string) *RedisKeyValueStateStoreBuilder {
	return &RedisKeyValueStateStoreBuilder{
		name:     name,
		isGlobal: isGlobal,
		address:  address,
	}
}

func (b *RedisKeyValueStateStoreBuilder) Build() (StateStore, error) {
	if b.address == "" {
		return nil, fmt.Errorf("redis address should not be empty")
	}
	return newRedisKeyValueStateStore(b.name, b.isGlobal, b.address)
}

func (b *RedisKeyValueStateStoreBuilder) Name() string {
	return b.name
}
