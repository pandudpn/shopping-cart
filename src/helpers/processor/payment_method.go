package processor

import (
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/spf13/viper"
)

func (p *processor) GetAvailablePaymentMethod(cart *model.Cart) error {
	maxEwallet := viper.GetFloat64("payment.max.ewallet")
	minCreditCard := viper.GetFloat64("payment.min.creditcard")

	paymentMethods, err := p.paymentMethodRepo.FindEnabledPaymentMethod()
	if err != nil {
		return err
	}

	for idx, paymentMethod := range paymentMethods {
		if (paymentMethod.Category == model.Ewallet && cart.Total > maxEwallet) || (paymentMethod.Category == model.CreditCard && cart.Total < minCreditCard) {
			paymentMethods = append(paymentMethods[:idx], paymentMethods[idx+1:]...)
		}
	}

	cart.SetAvailablePaymentMethod(paymentMethods)
	return nil
}
