package cached

import (
	"net/http"

	"github.com/pandudpn/shopping-cart/src/utils"
)

func responseError(w http.ResponseWriter, statusCode int, systemCode, errorMessage string, err error) {
	utils.Error(statusCode, systemCode, errorMessage, err).MiddlewareJson(w)
}
