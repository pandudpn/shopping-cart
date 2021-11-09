package datastorefactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/app/adapter/dbc/redis"
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

type redisFactory struct{}

func (rf *redisFactory) Build(c container.Container, enableTx bool) (DataStoreInterface, error) {
	if value, found := c.Get(constant.REDIS); found {
		logger.Log.Debugf("redis found in container %v", value)

		return value, nil
	}

	conn := dbc.RedisConnection()

	var rdbc = &redis.RedisDb{DB: conn}
	c.Put(constant.REDIS, rdbc)

	return rdbc, nil
}
