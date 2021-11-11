package user

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type UserRepository struct {
	DB dbc.SqlDbc
}

func rowToUser(row *sql.Row) (*model.User, error) {
	user := model.User{}

	err := row.Scan(&user.Id, &user.Name, &user.Email, &user.Phone, &user.Password, &user.Enabled, &user.EmailVerifiedAt)
	if err != nil {
		return nil, err
	}

	return &user, nil
}
