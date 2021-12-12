package paymenthandler

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/formatted"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
	"github.com/spf13/viper"
	"github.com/xendit/xendit-go"
	"github.com/xendit/xendit-go/ewallet"
)

type xenditEwallet struct{}

func (e *xenditEwallet) Process(cart *model.Cart) (PaymentHandlerInterface, error) {
	xendit.Opt.SecretKey = viper.GetString("xendit.api.secretkey")

	var (
		channelProperties map[string]string
		baskets           = make([]xendit.EWalletBasketItem, 0)
	)
	switch cart.GetPaymentMethod().Key {
	case OVO:
		channelProperties = map[string]string{
			"mobile_number": cart.GetUser().Phone,
		}
	case DANA, LINKAJA, SHOPEE:
		channelProperties = map[string]string{
			"success_redirect_url": fmt.Sprintf(viper.GetString("application.url.redirect"), cart.GetOrderNumber()),
		}
	default:
		return nil, errors.New("payment_method_handler.not_implemented")
	}

	for _, cartProduct := range cart.GetProducts() {
		basket := xendit.EWalletBasketItem{
			ReferenceID: fmt.Sprintf("%d", cartProduct.Product.Id),
			Name:        cartProduct.Product.Name,
			Category:    cartProduct.Product.Category.Name,
			Currency:    IDR,
			Quantity:    cartProduct.Quantity,
			Price:       cartProduct.BasePrice,
			Type:        Product,
		}

		baskets = append(baskets, basket)
	}

	data := ewallet.CreateEWalletChargeParams{
		ReferenceID:       cart.GetOrderNumber(),
		Currency:          IDR,
		Amount:            cart.Total,
		CheckoutMethod:    OneTimePayment,
		ChannelCode:       strings.ToUpper(cart.GetPaymentMethod().Key),
		ChannelProperties: channelProperties,
		Basket:            baskets,
	}

	charge, chargeErr := ewallet.CreateEWalletCharge(&data)
	if chargeErr != nil {
		logger.Log.Errorf("error charge e-wallet %v", chargeErr)
		return nil, chargeErr
	}
	logger.Log.Debugf("charge e-wallet %v", charge)

	expired, _ := time.ParseDuration(viper.GetString("xendit.expired.ewallet"))

	orderPayment := model.NewOrderPayment()
	orderPayment.PaymentMethodId = *cart.PaymentMethodId
	orderPayment.TotalPayment = charge.ChargeAmount
	orderPayment.ExpiredAt = orderPayment.CreatedAt.Add(expired)
	orderPayment.SetPaymentMethod(cart.GetPaymentMethod())

	if charge.IsRedirectRequired {
		if strings.ToLower(charge.ChannelCode) == SHOPEE {
			orderPayment.DeepLink = formatted.String(charge.Actions["mobile_deeplink_checkout_url"])
			orderPayment.QrLink = formatted.String(charge.Actions["qr_checkout_string"])
		} else {
			orderPayment.RedirectLink = formatted.String(charge.Actions["mobile_web_checkout_url"])
		}
	}

	return orderPayment, nil
}
