package checkoutpresenter

import (
	"sort"

	"github.com/pandudpn/shopping-cart/src/domain/model"
)

func categoryPaymentMethods() map[int]map[string]interface{} {
	return map[int]map[string]interface{}{
		0: {
			"category":       model.Ewallet,
			"paymentMethods": make([]*paymentMethodView, 0),
		},
		1: {
			"category":       model.VA,
			"paymentMethods": make([]*paymentMethodView, 0),
		},
		2: {
			"category":       model.CreditCard,
			"paymentMethods": make([]*paymentMethodView, 0),
		},
	}
}

func createAvailablePaymentMethods(availablePaymentMethods []*model.PaymentMethod) []*availablePaymentMethodView {
	categoryPaymentMethods := categoryPaymentMethods()
	availablePayments := make([]*availablePaymentMethodView, 0)

	keyCategory := make([]int, len(categoryPaymentMethods))
	i := 0
	for key := range categoryPaymentMethods {
		keyCategory[i] = key
		i++
	}
	sort.Ints(keyCategory)

	for key := range keyCategory {
		paymentMethods := make([]*paymentMethodView, 0)
		var category string

		for _, paymentMethod := range availablePaymentMethods {
			if categoryPaymentMethods[key]["category"] == paymentMethod.Category {
				paymentMethodView := createPaymentMethod(paymentMethod)

				paymentMethods = append(paymentMethods, paymentMethodView)
				category = paymentMethod.GetCategory()
			}
		}

		if len(paymentMethods) > 0 {
			availablePaymentMethod := &availablePaymentMethodView{
				Category:       category,
				PaymentMethods: paymentMethods,
			}

			availablePayments = append(availablePayments, availablePaymentMethod)
		}
	}

	return availablePayments
}

func createPaymentMethod(paymentMethod *model.PaymentMethod) *paymentMethodView {
	return &paymentMethodView{
		Id:       paymentMethod.Id,
		Code:     paymentMethod.Code,
		Category: paymentMethod.Category,
		Name:     paymentMethod.Name,
		Image:    paymentMethod.Image,
		Formatted: &formattedPaymentMethodView{
			Category: paymentMethod.GetCategory(),
		},
	}
}
