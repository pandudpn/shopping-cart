package containerhelper

import (
	"github.com/pandudpn/shopping-cart/app/container"
	"github.com/pandudpn/shopping-cart/src/api/controller"
)

func GetUserController(c container.Container) (controller.UserControllerInterface, error) {
	key := "user"

	uci, err := c.BuildController(key)
	if err != nil {
		return nil, err
	}

	return uci.(controller.UserControllerInterface), nil
}
