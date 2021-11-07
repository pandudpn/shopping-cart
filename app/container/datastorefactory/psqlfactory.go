package datastorefactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	sqlDbc "github.com/pandudpn/shopping-cart/app/adapter/dbc/sql"
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
)

type psqlFactory struct{}

func (pf *psqlFactory) Build(c container.Container, enableTx bool) (DataStoreInterface, error) {
	if !enableTx {
		if value, found := c.Get(constant.PSQL); found {
			return value, nil
		}
	}

	db := dbc.DatabaseConnection()
	var sdbc dbc.SqlDbc

	if enableTx {
		tx, err := db.Begin()
		if err != nil {
			return nil, err
		}

		sdbc = &sqlDbc.SqlTx{DB: tx}
	} else {
		sdbc = &sqlDbc.SqlDb{DB: db}

		c.Put(constant.PSQL, sdbc)
	}

	return sdbc, nil
}
