package redis

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisDb struct {
	DB *redis.Client
}

func (rd *RedisDb) Get(ctx context.Context, key string) *redis.StringCmd {
	return rd.DB.Get(ctx, key)
}

func (rd *RedisDb) Set(ctx context.Context, key string, value interface{}, expiration time.Duration) *redis.StatusCmd {
	return rd.DB.Set(ctx, key, value, expiration)
}

func (rd *RedisDb) Del(ctx context.Context, key ...string) *redis.IntCmd {
	return rd.DB.Del(ctx, key...)
}

func (rd *RedisDb) GetDel(ctx context.Context, key string) *redis.StringCmd {
	return rd.DB.GetDel(ctx, key)
}
