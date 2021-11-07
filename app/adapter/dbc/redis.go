package dbc

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
)

func RedisConnection() *redis.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	dsn := viper.GetString("redis.dsn")
	dbNum := viper.GetInt("redis.dbnum")
	pw := viper.GetString("redis.password")

	opts := &redis.Options{
		Addr:     dsn,
		Password: pw,
		DB:       dbNum,
	}

	conn := redis.NewClient(opts)

	err := conn.Ping(ctx).Err()
	if err != nil {
		panic(err)
	}

	return conn
}
