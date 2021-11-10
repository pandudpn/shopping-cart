package usercontroller

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/pandudpn/shopping-cart/src/api/presenter/userpresenter"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (uc *UserController) RegisterHandler(e echo.Context) error {
	var (
		req     = e.Request()
		ctx     = req.Context()
		payload = &model.RequestRegister{}
	)

	if err := e.Bind(&payload); err != nil {
		logger.Log.Errorf("error parsing payload register %v", err)
		err = errors.New("body.payload")
		return userpresenter.ResponseRegister(nil, err).JSON(e)
	}

	return uc.UserUseCase.RegisterUser(ctx, payload).JSON(e)
}
