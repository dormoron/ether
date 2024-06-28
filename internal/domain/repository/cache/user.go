package cache

import (
	"context"
	"encoding/json"
	"ether/internal/domain/model"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type UserCache interface {
	Get(ctx context.Context, id uint) (model.User, error)
	Set(ctx context.Context, u model.User) error
	key(id uint) string
}

type RedisUserCache struct {
	cmd        redis.Cmdable
	expiration time.Duration
}

func NewUserCache(cmd redis.Cmdable) UserCache {
	return &RedisUserCache{
		cmd:        cmd,
		expiration: time.Minute * 15,
	}
}

func (cache *RedisUserCache) Get(ctx context.Context, id uint) (model.User, error) {
	val, err := cache.cmd.Get(ctx, cache.key(id)).Bytes()
	if err != nil {
		return model.User{}, err
	}
	var u model.User
	err = json.Unmarshal(val, &u)
	return u, err
}

func (cache *RedisUserCache) Set(ctx context.Context, u model.User) error {
	val, err := json.Marshal(u)
	if err != nil {
		return err
	}
	return cache.cmd.Set(ctx, cache.key(u.Id), val, cache.expiration).Err()
}

func (cache *RedisUserCache) key(id uint) string {
	return fmt.Sprintf("user:info:%d", id)
}
