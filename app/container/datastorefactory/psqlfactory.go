package datastorefactory

import (
	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	sqlDbc "github.com/pandudpn/shopping-cart/app/adapter/dbc/sql"
	"github.com/pandudpn/shopping-cart/app/constant"
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

type psqlFactory struct{}

func (pf *psqlFactory) Build(c container.Container, enableTx bool) (DataStoreInterface, error) {
	if !enableTx {
		if value, found := c.Get(constant.PSQL); found {
			logger.Log.Debug("psql found in container")
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
		logger.Log.Debug("psql.Tx connected")
		sdbc = &sqlDbc.SqlTx{DB: tx}
	} else {
		sdbc = &sqlDbc.SqlDb{DB: db}
		logger.Log.Debug("psql.DB connected")
		c.Put(constant.PSQL, sdbc)
	}

	return sdbc, nil
}
