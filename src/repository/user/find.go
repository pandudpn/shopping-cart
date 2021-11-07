package user

import "github.com/pandudpn/shopping-cart/src/domain/model"

const (
	QUERY_USER_BY_EMAIL = "select id, name, email from users where email=$1"
	QUERY_USER_BY_PHONE = "select id, name, email from users where phone=$1"
)

func (ur *UserRepository) FindByEmail(email string) (*model.User, error) {
	row := ur.DB.QueryRow(QUERY_USER_BY_EMAIL, email)

	return rowToUser(row)
}

func (ur *UserRepository) FindByPhone(phone string) (*model.User, error) {
	rows, err := ur.DB.Query(QUERY_USER_BY_PHONE, phone)
	if err != nil {
		return nil, err
	}

	return retrieveUser(rows)
}
