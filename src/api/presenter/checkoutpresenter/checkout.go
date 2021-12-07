package checkoutpresenter

import (
	"net/http"

	"github.com/pandudpn/shopping-cart/src/utils"
)

var (
	errGlobal string = "Terjadi kesalahan pada server, silakan coba beberapa saat lagi"

	getCheckoutSuccess    = "checkout.get.success"
	getCheckoutFailed     = "checkout.get.failed"
	updateCheckoutSuccess = "checkout.update.success"
	updateCheckoutFailed  = "checkout.update.failed"
	cartActive            = "checkout.active.failed"

	keyRequired = "header.key.required"
	queryFind   = "query.find.error"
	queryInsert = "query.insert.error"
	queryUpdate = "query.update.failed"

	message = map[string]string{
		queryFind:             errGlobal,
		queryInsert:           errGlobal,
		queryUpdate:           errGlobal,
		cartActive:            errGlobal,
		updateCheckoutSuccess: "Berhasil diubah",
		updateCheckoutFailed:  "Data gagal diubah",
		getCheckoutFailed:     "Gagal mengambil keranjang belanja anda",
		keyRequired:           "Kunci keranjang tidak ditemukan",
	}

	systemCode = map[string]string{
		cartActive:            "44",
		getCheckoutSuccess:    "50",
		updateCheckoutSuccess: "51",
		getCheckoutFailed:     "52",
		updateCheckoutFailed:  "53",
		queryFind:             "81",
		queryInsert:           "82",
		queryUpdate:           "83",
		keyRequired:           "84",
	}

	statusCode = map[string]int{
		keyRequired:           http.StatusBadRequest,
		getCheckoutSuccess:    http.StatusOK,
		getCheckoutFailed:     http.StatusBadRequest,
		updateCheckoutSuccess: http.StatusOK,
		updateCheckoutFailed:  http.StatusBadRequest,
		queryFind:             http.StatusInternalServerError,
		queryInsert:           http.StatusInternalServerError,
		queryUpdate:           http.StatusInternalServerError,
		keyRequired:           http.StatusBadRequest,
		cartActive:            http.StatusInternalServerError,
	}
)

func ResponseCheckout(isCheckoutProgress bool, value interface{}, err error) utils.ResponseInterface {
	if err != nil {
		var errMessage = err.Error()

		return utils.Error(statusCode[errMessage], systemCode[errMessage], message[errMessage], err)
	}

	if isCheckoutProgress {
		return utils.Success(statusCode[updateCheckoutSuccess], systemCode[updateCheckoutSuccess], value)
	}

	return utils.Success(statusCode[getCheckoutSuccess], systemCode[getCheckoutSuccess], value)
}
