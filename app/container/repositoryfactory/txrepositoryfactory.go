package repositoryfactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/datastorefactory"
	"github.com/pandudpn/shopping-cart/src/repository/sqltx"
)

type txRepositoryFactory struct{}

func (trf *txRepositoryFactory) Build(c container.Container) (RepositoryFactoryInterface, error) {
	code := constant.PSQL

	dsfi, err := datastorefactory.GetDataStoreFbMap(code).Build(c, constant.ENABLETX)
	if err != nil {
		return nil, err
	}

	rus := dsfi.(dbc.SqlDbc)
	ur := sqltx.TxRepository{DB: rus}

	return &ur, nil
}
