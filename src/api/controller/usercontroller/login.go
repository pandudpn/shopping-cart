package usercontroller

import (
	"github.com/labstack/echo"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

func (uc *UserController) LoginHandler(e echo.Context) error {
	var (
		req  = e.Request()
		ctx  = req.Context()
		user = &model.User{}
	)
	user.Email = "pandu@pandudpn.id"

	return uc.UserUseCase.LoginUser(ctx, user.Email).JSON(e)
}
