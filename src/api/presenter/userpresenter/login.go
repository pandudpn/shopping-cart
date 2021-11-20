package userpresenter

import (
	"net/http"

	"github.com/pandudpn/shopping-cart/src/utils"
)

var (
	errGlobal            string = "Terjadi kesalahan pada server, silakan coba beberapa saat lagi"
	userEmailNotFound    string = "user.email.not_found"
	userPhoneNotFound    string = "user.phone.not_found"
	userNotActive        string = "user.not_active"
	userPasswordNotMatch string = "user.password.not_match"
	queryError           string = "query.find.error"
	insertError          string = "query.insert.error"
	bodyPayload          string = "body.payload"
	errCreateSession     string = "session.create.error"
	loginSuccess         string = "login.success"

	message = map[string]string{
		userEmailNotFound:    "User tidak ditemukan",
		userPhoneNotFound:    "User tidak ditemukan",
		userNotActive:        "User belum aktif",
		userPasswordNotMatch: "Email atau password tidak sesuai",
		queryError:           errGlobal,
		insertError:          errGlobal,
		errCreateSession:     errGlobal,
		bodyPayload:          "Permintaan kamu tidak lengkap",

		registerSuccess:   "Register berhasil, silahkan login sekarang juga",
		emailAlreadyTaken: "Email sudah digunakan",
		phoneAlreadyTaken: "Nomer telepon sudah digunakan",
		phoneInvalid:      "Nomer telepon tidak sesuai, pastikan nomer telepon sesuai",
	}

	systemCode = map[string]string{
		loginSuccess:         "10",
		userEmailNotFound:    "12",
		userPhoneNotFound:    "13",
		userNotActive:        "14",
		userPasswordNotMatch: "15",

		registerSuccess:   "20",
		emailAlreadyTaken: "22",
		phoneAlreadyTaken: "23",
		phoneInvalid:      "24",

		bodyPayload:      "80",
		queryError:       "81",
		insertError:      "82",
		errCreateSession: "89",
	}

	statusCode = map[string]int{
		loginSuccess:         http.StatusOK,
		userEmailNotFound:    http.StatusNotFound,
		userPhoneNotFound:    http.StatusNotFound,
		userNotActive:        http.StatusUnprocessableEntity,
		userPasswordNotMatch: http.StatusBadRequest,
		registerSuccess:      http.StatusCreated,
		emailAlreadyTaken:    http.StatusBadRequest,
		phoneAlreadyTaken:    http.StatusBadRequest,
		phoneInvalid:         http.StatusBadRequest,
		queryError:           http.StatusInternalServerError,
		insertError:          http.StatusInternalServerError,
		errCreateSession:     http.StatusInternalServerError,
		bodyPayload:          http.StatusBadRequest,
	}
)

func ResponseLogin(value interface{}, err error) utils.ResponseInterface {
	if err != nil {
		errString := err.Error()
		return utils.Error(statusCode[errString], systemCode[errString], message[errString], err)
	}

	return utils.Success(statusCode[loginSuccess], systemCode[loginSuccess], value)
}
