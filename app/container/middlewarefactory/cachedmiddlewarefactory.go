package middlewarefactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/datastorefactory"
	"github.com/pandudpn/shopping-cart/src/api/middleware/cached"
)

type redisMiddlewareFactory struct{}

func (rmf *redisMiddlewareFactory) Build(c container.Container) (MiddlewareFactoryInterface, error) {
	dsrf, err := datastorefactory.GetDataStoreFbMap(constant.REDIS).Build(c, !constant.ENABLETX)
	if err != nil {
		return nil, err
	}

	dsr := dsrf.(dbc.RDbc)
	rm := cached.RedisMiddleware{RedisDb: dsr}

	return &rm, nil
}
