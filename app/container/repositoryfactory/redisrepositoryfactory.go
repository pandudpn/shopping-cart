package repositoryfactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/datastorefactory"
	"github.com/pandudpn/shopping-cart/src/repository/rdb"
)

type redisRepositoryFactory struct{}

func (rrf *redisRepositoryFactory) Build(c container.Container) (RepositoryFactoryInterface, error) {
	code := constant.REDIS

	rfi, err := datastorefactory.GetDataStoreFbMap(code).Build(c, !constant.ENABLETX)
	if err != nil {
		return nil, err
	}

	rdbc := rfi.(dbc.RDbc)
	rr := rdb.RedisRepository{RDb: rdbc}

	return &rr, nil
}
