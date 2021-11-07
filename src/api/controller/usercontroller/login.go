package usercontroller

import (
	"net/http"

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

	user, err := uc.UserUseCase.LoginUser(ctx, user)
	if err != nil {
		return e.JSON(http.StatusFound, err)
	}

	return e.JSON(http.StatusOK, user)
}
