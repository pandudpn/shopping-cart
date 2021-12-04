package repositoryfactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/datastorefactory"
	"github.com/pandudpn/shopping-cart/src/repository/courierrepo"
)

type courierRepositoryFactory struct{}

func (crf *courierRepositoryFactory) Build(c container.Container) (RepositoryFactoryInterface, error) {
	dsfi, err := datastorefactory.GetDataStoreFbMap(constant.PSQL).Build(c, !constant.ENABLETX)
	if err != nil {
		return nil, err
	}

	db := dsfi.(dbc.SqlDbc)
	cr := courierrepo.CourierRepository{DB: db}

	return &cr, nil
}
