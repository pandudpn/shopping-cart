package usercontroller

import "github.com/pandudpn/shopping-cart/src/domain/usecase"

type UserController struct {
	UserUseCase usecase.UserUseCaseInterface
}
