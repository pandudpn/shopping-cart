package callbackcontroller

import (
	"errors"

	"github.com/labstack/echo"
	"github.com/pandudpn/shopping-cart/src/api/presenter/callbackpresenter"
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (cc *CallbackController) XenditEWalletHandler(e echo.Context) error {
	var (
		r       = e.Request()
		ctx     = r.Context()
		payload = &model.EWalletPaymentNotification{}
	)

	if err := e.Bind(&payload); err != nil {
		logger.Log.Errorf("error parsing payload callback xendit-ewallet %v", err)
		err = errors.New("body.payload")
		return callbackpresenter.ResponseCallback(nil, err).JSON(e)
	}

	logger.Log.Infof("payload %v", payload)

	return cc.CallbackUseCase.CallbackEWallet(ctx, payload).JSON(e)
}
