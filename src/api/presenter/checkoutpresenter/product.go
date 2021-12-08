package checkoutpresenter

import (
	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/formatted"
)

func createCartProductView(cart *model.Cart) ([]*cartProductView, float64) {
	cartProductsView := make([]*cartProductView, 0)
	var totalProductsPrice float64

	for _, cartProduct := range cart.GetProducts() {
		cartProductView := &cartProductView{
			Id:         cartProduct.Id,
			Quantity:   cartProduct.Quantity,
			Price:      int(cartProduct.BasePrice),
			TotalPrice: int(cartProduct.TotalPrice),
			Product:    createProductView(cartProduct.GetProduct()),
			Formatted: &formattedCartProductView{
				Price:      formatted.IndonesiaCurrrency(cartProduct.BasePrice),
				TotalPrice: formatted.IndonesiaCurrrency(cartProduct.TotalPrice),
			},
		}

		cartProductsView = append(cartProductsView, cartProductView)
		totalProductsPrice += cartProduct.TotalPrice
	}

	return cartProductsView, totalProductsPrice
}

func createProductView(product *model.Product) *productView {
	productView := &productView{
		Id:              product.Id,
		Name:            product.Name,
		Slug:            product.Slug,
		Price:           int(product.Price),
		DiscountedPrice: int(product.DiscountedPrice),
		Qty:             product.GetQuantity(),
		Formatted: &formattedProductView{
			Price:           formatted.IndonesiaCurrrency(product.Price),
			DiscountedPrice: formatted.IndonesiaCurrrency(product.DiscountedPrice),
		},
	}

	if len(product.GetImages()) > 0 {
		images := make([]*imageView, 0)
		for _, image := range product.GetImages() {
			img := &imageView{
				Url: image.GetImage().GetFile(),
			}

			images = append(images, img)
		}

		productView.Images = images
	}

	return productView
}
