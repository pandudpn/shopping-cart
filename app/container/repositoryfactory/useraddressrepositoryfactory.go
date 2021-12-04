package repositoryfactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/app/container/datastorefactory"
	"github.com/pandudpn/shopping-cart/src/repository/useraddressrepo"
)

type userAddressRepositoryFactory struct{}

func (uarf *userAddressRepositoryFactory) Build(c container.Container) (RepositoryFactoryInterface, error) {
	dsfi, err := datastorefactory.GetDataStoreFbMap(constant.PSQL).Build(c, !constant.ENABLETX)
	if err != nil {
		return nil, err
	}

	db := dsfi.(dbc.SqlDbc)
	uar := useraddressrepo.UserAddressRepository{
		DB: db,
	}

	return &uar, nil
}
