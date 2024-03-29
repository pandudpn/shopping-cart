package cached

import (
	"net/http"

	"github.com/pandudpn/shopping-cart/src/utils"
)

func responseError(w http.ResponseWriter, statusCode int, systemCode, errorMessage string, err error) {
	utils.Error(statusCode, systemCode, errorMessage, err).MiddlewareJson(w)
}

func responseSuccess(w http.ResponseWriter, value interface{}) {
	utils.Success(http.StatusOK, CachedData, value).MiddlewareJson(w)
}
