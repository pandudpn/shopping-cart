package repository

import "github.com/pandudpn/shopping-cart/src/domain/model"

type UserRepositoryInterface interface {
	FindByEmail(email string) (*model.User, error)
}
