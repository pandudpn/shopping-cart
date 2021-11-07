package middlewarefactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/datastorefactory"
	"github.com/pandudpn/shopping-cart/src/api/middleware/cached"
)

type cachedMiddlewareFactory struct{}

func (cmf *cachedMiddlewareFactory) Build(c container.Container) (MiddlewareFactoryInterface, error) {
	dsrf, err := datastorefactory.GetDataStoreFbMap(constant.REDIS).Build(c, !constant.ENABLETX)
	if err != nil {
		return nil, err
	}

	dsr := dsrf.(dbc.RDbc)
	cm := cached.SessionMiddleware{RedisDb: dsr}

	return &cm, nil
}
