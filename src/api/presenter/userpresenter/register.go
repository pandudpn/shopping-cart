package userpresenter

import (
	"github.com/pandudpn/shopping-cart/src/utils"
)

var (
	registerSuccess   string = "register.success"
	emailAlreadyTaken string = "user.email.taken"
	phoneAlreadyTaken string = "user.phone.taken"
	phoneInvalid      string = "phone.invalid"
)

func ResponseRegister(value interface{}, err error) utils.ResponseInterface {
	var valueString = "register.success"

	if err != nil {
		errMsg := err.Error()
		return utils.Error(statusCode[errMsg], systemCode[errMsg], message[errMsg], err)
	}
	if v, f := value.(string); f {
		valueString = v
		value = message[v]
	}

	return utils.Success(statusCode[valueString], systemCode[valueString], value)
}
