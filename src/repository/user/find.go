package user

import (
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

const (
	QUERY_USER_SELECT   = `select id, name, email, phone, password, enabled, email_verified_at from "user" `
	QUERY_USER_BY_EMAIL = QUERY_USER_SELECT + "where email=$1"
	QUERY_USER_BY_PHONE = QUERY_USER_SELECT + "where phone=$1"
	QUERY_USER_BY_ID    = QUERY_USER_SELECT + "where id=$1"
)

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	row := ur.DB.QueryRow(QUERY_USER_BY_EMAIL, email)

	return rowToUser(row)
}

func (ur *UserRepository) FindByPhone(phone string) (*model.User, error) {
	row := ur.DB.QueryRow(QUERY_USER_BY_PHONE, phone)

	return rowToUser(row)
}

func (ur *UserRepository) FindById(id int) (*model.User, error) {
	row := ur.DB.QueryRow(QUERY_USER_BY_ID, id)

	return rowToUser(row)
}
