package datastorefactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/app/adapter/dbc/redis"
	"github.com/pandudpn/shopping-cart/app/container"
)

type redisFactory struct{}

func (rf *redisFactory) Build(c container.Container, enableTx bool) (DataStoreInterface, error) {
	conn := dbc.RedisConnection()

	var rdbc = &redis.RedisDb{DB: conn}

	return rdbc, nil
}
