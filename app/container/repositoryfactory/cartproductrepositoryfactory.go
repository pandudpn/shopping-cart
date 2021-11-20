package repositoryfactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/datastorefactory"
	"github.com/pandudpn/shopping-cart/src/repository/cartproductrepo"
)

type cartProductRepositoryFactory struct{}

func (cprf *cartProductRepositoryFactory) Build(c container.Container) (RepositoryFactoryInterface, error) {
	code := constant.PSQL

	dsfi, err := datastorefactory.GetDataStoreFbMap(code).Build(c, !constant.ENABLETX)
	if err != nil {
		return nil, err
	}

	db := dsfi.(dbc.SqlDbc)
	cr := cartproductrepo.CartProductRepository{DB: db}

	return &cr, nil
}
