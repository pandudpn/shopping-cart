package checkoutcontroller

import "github.com/pandudpn/shopping-cart/src/domain/usecase"

type CheckoutController struct {
	CheckoutUseCase usecase.CheckoutUseCaseInterface
}
