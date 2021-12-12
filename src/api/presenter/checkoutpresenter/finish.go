package checkoutpresenter

import (
	"net/http"

	"github.com/pandudpn/shopping-cart/src/utils"
)

func FinishCheckout(value interface{}, err error) utils.ResponseInterface {
	if err != nil {
		return utils.Error(http.StatusBadRequest, "2000", err.Error(), err)
	}

	return utils.Success(http.StatusCreated, "1000", value)
}
