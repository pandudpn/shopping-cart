package cartpresenter

import (
	"net/http"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils"
)

type cartView struct {
	Id                   int                `json:"id"`
	Key                  string             `json:"key"`
	IsActive             bool               `json:"isActive"`
	DeliveryCost         float64            `json:"deliveryCost"`
	DeliveryCostDiscount *float64           `json:"deliveryCostDiscount"`
	TotalProductPrice    float64            `json:"totalProductPrice"`
	TotalDeliveryCost    float64            `json:"totalDeliveryCost"`
	Customer             *userView          `json:"customer"`
	Products             []*cartProductView `json:"products"`
}

type userView struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phoneNumber"`
}

type cartProductView struct {
	Id         int          `json:"id"`
	Quantity   int          `json:"quantity"`
	Price      int          `json:"price"`
	TotalPrice int          `json:"totalPrice"`
	Product    *productView `json:"product"`
}

type productView struct {
	Id              int          `json:"id"`
	Name            string       `json:"name"`
	Slug            string       `json:"slug"`
	Price           int          `json:"price"`
	DiscountedPrice int          `json:"discountedPrice"`
	Qty             int          `json:"qty"`
	Images          []*imageView `json:"images"`
}

type imageView struct {
	Url string `json:"url"`
}

var (
	errGlobal string = "Terjadi kesalahan pada server, silakan coba beberapa saat lagi"

	getCartSuccess   = "cart.success"
	getCartFailed    = "cart.failed"
	addToCartSuccess = "cart.add_product"
	cartActive       = "cart.active.failed"

	productNotFound = "cart.product.not_found"
	productQuantity = "cart.product.quantity"

	bodyPayload = "body.payload"
	queryFind   = "query.find.error"
	queryInsert = "query.insert.error"

	message = map[string]string{
		queryFind:       errGlobal,
		queryInsert:     errGlobal,
		cartActive:      errGlobal,
		productNotFound: "Produk tidak ditemukan",
		productQuantity: "Stok produk tidak mencukupi",
		getCartFailed:   "Gagal mengambil keranjang belanja anda",
		bodyPayload:     "Permintaan kamu tidak lengkap",
	}

	systemCode = map[string]string{
		productNotFound:  "32",
		productQuantity:  "35",
		getCartSuccess:   "40",
		addToCartSuccess: "41",
		getCartFailed:    "43",
		cartActive:       "44",
		bodyPayload:      "80",
		queryFind:        "81",
		queryInsert:      "82",
	}

	statusCode = map[string]int{
		productNotFound:  http.StatusNotFound,
		productQuantity:  http.StatusBadRequest,
		getCartSuccess:   http.StatusOK,
		addToCartSuccess: http.StatusCreated,
		getCartFailed:    http.StatusBadRequest,
		queryFind:        http.StatusInternalServerError,
		queryInsert:      http.StatusInternalServerError,
		bodyPayload:      http.StatusBadRequest,
		cartActive:       http.StatusInternalServerError,
	}
)

func ResponseCart(isAddToCart bool, value interface{}, err error) utils.ResponseInterface {
	if err != nil {
		var errMessage = err.Error()

		return utils.Error(statusCode[errMessage], systemCode[errMessage], message[errMessage], err)
	}

	var (
		resStatusCode = statusCode[getCartSuccess]
		resSystemCode = systemCode[getCartSuccess]
	)

	if isAddToCart {
		resStatusCode = statusCode[addToCartSuccess]
		resSystemCode = systemCode[addToCartSuccess]
	}

	if cart, ok := value.(*model.Cart); ok {
		res := createViewCart(cart)

		return utils.Success(resStatusCode, resSystemCode, res)
	}

	return utils.Success(statusCode[getCartSuccess], systemCode[getCartSuccess], value)
}

func createViewCart(cart *model.Cart) *cartView {
	products, priceProduct := createCartProductView(cart)

	res := &cartView{
		Id:                cart.Id,
		Key:               cart.Key,
		IsActive:          cart.IsActive,
		TotalProductPrice: priceProduct,
		Products:          products,
	}

	return res
}

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
