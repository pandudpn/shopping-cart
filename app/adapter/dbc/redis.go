package dbc

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
	"github.com/spf13/viper"
)

func RedisConnection() *redis.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
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
		logger.Log.Error(err)
		panic(err)
	}

	return conn
}
