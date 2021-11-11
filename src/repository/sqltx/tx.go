package sqltx

import (
	"errors"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/app/adapter/dbc/sql"
)

type TxRepository struct {
	DB dbc.SqlDbc
}

func (tr *TxRepository) TxEnd(txFunc func() error) error {
	if _, isTx := tr.DB.(*sql.SqlTx); !isTx {
		err := errors.New("connection not for transaction")
		return err
	}

	return tr.DB.TxEnd(txFunc)
}
