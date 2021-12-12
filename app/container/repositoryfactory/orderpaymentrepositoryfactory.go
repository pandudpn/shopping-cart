package repositoryfactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/datastorefactory"
	"github.com/pandudpn/shopping-cart/src/repository/orderpaymentrepo"
)

type orderPaymentRepositoryFactory struct{}

func (oprf *orderPaymentRepositoryFactory) Build(c container.Container) (RepositoryFactoryInterface, error) {
	code := constant.PSQL

	dsfi, err := datastorefactory.GetDataStoreFbMap(code).Build(c, !constant.ENABLETX)
	if err != nil {
		return nil, err
	}

	db := dsfi.(dbc.SqlDbc)
	opri := orderpaymentrepo.OrderPaymentRepository{DB: db}

	return &opri, nil
}
