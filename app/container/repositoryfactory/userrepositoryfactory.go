package repositoryfactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/datastorefactory"
	"github.com/pandudpn/shopping-cart/src/repository/user"
)

type userRepositoryFactory struct{}

func (urf *userRepositoryFactory) Build(c container.Container, enabledTx bool) (RepositoryFactoryInterface, error) {
	code := constant.PSQL

	dsfi, err := datastorefactory.GetDataStoreFbMap(code).Build(c, enabledTx)
	if err != nil {
		return nil, err
	}

	rus := dsfi.(dbc.SqlDbc)
	ur := user.UserRepository{DB: rus}

	return &ur, nil
}
