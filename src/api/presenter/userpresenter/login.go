package userpresenter

import (
	"net/http"

	"github.com/pandudpn/shopping-cart/src/utils"
)

var (
	errGlobal         string = "Terjadi kesalahan pada server, silakan coba beberapa saat lagi"
	userEmailNotFound string = "user.email.not_found"
	userPhoneNotFound string = "user.phone.not_found"
	userNotActive     string = "user.not_active"
	queryError        string = "query.find.error"
	errCreateSession  string = "session.create.error"

	message = map[string]string{
		userEmailNotFound: "User tidak ditemukan",
		userPhoneNotFound: "User tidak ditemukan",
		userNotActive:     "User belum aktif",
		queryError:        errGlobal,
		errCreateSession:  errGlobal,
	}

	systemCode = map[string]string{
		"login.success":   "10",
		userEmailNotFound: "12",
		userPhoneNotFound: "13",
		userNotActive:     "14",
		queryError:        "81",
		errCreateSession:  "89",
	}

	statusCode = map[string]int{
		"login.success":   http.StatusOK,
		userEmailNotFound: http.StatusNotFound,
		userPhoneNotFound: http.StatusNotFound,
		userNotActive:     http.StatusUnprocessableEntity,
		queryError:        http.StatusInternalServerError,
		errCreateSession:  http.StatusInternalServerError,
	}
)

func ResponseLogin(value interface{}, err error) utils.ResponseInterface {
	if err != nil {
		errString := err.Error()
		return utils.Error(statusCode[errString], systemCode[errString], message[errString], err)
	}

	return utils.Success(statusCode["login.success"], systemCode["login.success"], value)
}
