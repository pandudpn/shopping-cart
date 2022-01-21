package callbackpresenter

import (
	"net/http"

	"github.com/pandudpn/shopping-cart/src/utils"
)

var (
	callbackSuccess  = "callback.success"
	orderNotfound    = "order.notfound"
	databaseRelation = "database.relation.error"

	createDeliveryError = "courier.create_delivery.error"

	bodyPayload = "body.payload"

	message = map[string]string{
		callbackSuccess:     "Notifikasi pembayaran sukses",
		orderNotfound:       "Order Number tidak ditemukan",
		databaseRelation:    "Gagal mengambil data product, payment atau delivery",
		bodyPayload:         "Permintaan kamu tidak lengkap",
		createDeliveryError: "Terjadi kesalahan saat membuat delivery courier",
	}

	systemCode = map[string]string{
		callbackSuccess:     "71",
		databaseRelation:    "73",
		createDeliveryError: "74",
		orderNotfound:       "75",
		bodyPayload:         "80",
	}

	statusCode = map[string]int{
		callbackSuccess:     http.StatusOK,
		databaseRelation:    http.StatusInternalServerError,
		createDeliveryError: http.StatusInternalServerError,
		orderNotfound:       http.StatusNotFound,
		bodyPayload:         http.StatusBadRequest,
	}
)

func ResponseCallback(val interface{}, err error) utils.ResponseInterface {
	if err != nil {
		msg := err.Error()

		return utils.Error(statusCode[msg], systemCode[msg], message[msg], err)
	}

	return utils.Success(statusCode[callbackSuccess], systemCode[callbackSuccess], val)
}
