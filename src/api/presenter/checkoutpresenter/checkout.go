package checkoutpresenter

import (
	"net/http"
	"sync"
	"time"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils"
)

type checkoutView struct {
	Id                      int                           `json:"id"`
	Key                     string                        `json:"key"`
	TotalProductsPrice      float64                       `json:"totalProductsPrice"`
	DeliveryCost            float64                       `json:"deliveryCost"`
	DeliveryCostDiscount    float64                       `json:"deliveryCostDiscount"`
	TotalDeliveryCost       float64                       `json:"totalDeliveryCost"`
	TotalPayment            float64                       `json:"totalPayment"`
	Customer                *customerView                 `json:"customer"`
	DeliveryAddress         *deliveryAddressView          `json:"deliveryAddress"`
	Courier                 *courierView                  `json:"courier"`
	PaymentMethod           *paymentMethodView            `json:"paymentMethod"`
	Products                []*cartProductView            `json:"products"`
	AvailableCouriers       []*availableCourierView       `json:"availableCouriers"`
	AvailablePaymentMethods []*availablePaymentMethodView `json:"availablePaymentMethods"`
}

type customerView struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	PhoneNumber string `json:"phoneNumber"`
}

type deliveryAddressView struct {
	Id           int         `json:"id"`
	ReceiverName string      `json:"receiverName"`
	PhoneNumber  string      `json:"phoneNumber"`
	Address      string      `json:"address"`
	PostCode     string      `json:"postCode"`
	Lat          *float64    `json:"lat"`
	Long         *float64    `json:"long"`
	Province     *regionView `json:"province"`
	City         *regionView `json:"city"`
	District     *regionView `json:"district"`
	Area         *regionView `json:"area"`
}

type regionView struct {
	Id   int      `json:"id"`
	Name string   `json:"name"`
	Lat  *float64 `json:"lat"`
	Long *float64 `json:"long"`
}

type availableCourierView struct {
	Category string         `json:"category"`
	Couriers []*courierView `json:"couriers"`
}

type courierView struct {
	Id              int                   `json:"id"`
	Code            string                `json:"code"`
	Name            string                `json:"name"`
	Category        string                `json:"category"`
	Price           float64               `json:"price"`
	DiscountedPrice float64               `json:"discountedPrice"`
	InsurancePrice  float64               `json:"insurancePrice"`
	TotalPrice      float64               `json:"totalPrice"`
	Image           *string               `json:"image"`
	IsEligible      bool                  `json:"isEligible"`
	Label           string                `json:"label"`
	Formatted       *formattedCourierView `json:"formatted"`
}

type formattedCourierView struct {
	Price           string `json:"price"`
	DiscountedPrice string `json:"discountedPrice"`
	InsurancePrice  string `json:"insurancePrice"`
	TotalPrice      string `json:"totalPrice"`
	Category        string `json:"category"`
}

type availablePaymentMethodView struct {
	Category       string               `json:"category"`
	PaymentMethods []*paymentMethodView `json:"paymentMethods"`
}

type paymentMethodView struct {
	Id        int                         `json:"id"`
	Code      string                      `json:"code"`
	Category  string                      `json:"category"`
	Name      string                      `json:"name"`
	Image     *string                     `json:"image"`
	Formatted *formattedPaymentMethodView `json:"formatted"`
}

type formattedPaymentMethodView struct {
	Category string `json:"category"`
}

type cartProductView struct {
	Id         int                       `json:"id"`
	Quantity   int                       `json:"quantity"`
	Price      int                       `json:"price"`
	TotalPrice int                       `json:"totalPrice"`
	Formatted  *formattedCartProductView `json:"formatted"`
	Product    *productView              `json:"product"`
}

type formattedCartProductView struct {
	Price      string `json:"price"`
	TotalPrice string `json:"totalPrice"`
}

type productView struct {
	Id              int                   `json:"id"`
	Name            string                `json:"name"`
	Slug            string                `json:"slug"`
	Price           int                   `json:"price"`
	DiscountedPrice int                   `json:"discountedPrice"`
	Qty             int                   `json:"qty"`
	Images          []*imageView          `json:"images"`
	Formatted       *formattedProductView `json:"formatted"`
}

type imageView struct {
	Url string `json:"url"`
}

type formattedProductView struct {
	Price           string `json:"price"`
	DiscountedPrice string `json:"discountedPrice"`
}

var (
	timezone, _           = time.LoadLocation("Asia/Jakarta")
	datetimeFormat        = "2006-01-02T15:04:05-0700"
	errGlobal      string = "Terjadi kesalahan pada server, silakan coba beberapa saat lagi"

	getCheckoutSuccess    = "checkout.get.success"
	getCheckoutFailed     = "checkout.get.failed"
	updateCheckoutSuccess = "checkout.update.success"
	updateCheckoutFailed  = "checkout.update.failed"
	cartActive            = "checkout.active.failed"
	paymentMethod         = "cart.payment_method.not_found"
	finishCheckoutSuccess = "checkout.finish.success"
	finishCheckoutFailed  = "checkout.finish.failed"

	paymentError = "payment.create.error"

	deliveryAddressNotHave = "checkout.delivery_address.not_found"
	paymentMethodError     = "checkout.payment_method.error"
	courierUnavail         = "courier.not_avail"
	courierError           = "courier.error"
	courierNotFound        = "cart.courier.not_found"
	userNotFound           = "cart.user.not_found"
	userAddressNotFound    = "cart.user_address.not_found"
	cartNotYours           = "cart.not_yours"

	keyRequired = "header.key.required"
	queryFind   = "query.find.error"
	queryInsert = "query.insert.error"
	queryUpdate = "query.update.failed"
	bodyPayload = "body.payload"

	message = map[string]string{
		queryFind:              errGlobal,
		queryInsert:            errGlobal,
		queryUpdate:            errGlobal,
		cartActive:             errGlobal,
		updateCheckoutSuccess:  "Berhasil diubah",
		updateCheckoutFailed:   "Data gagal diubah",
		getCheckoutFailed:      "Gagal mengambil keranjang belanja anda",
		keyRequired:            "Kunci keranjang tidak ditemukan",
		deliveryAddressNotHave: "Tujuan pengiriman belum di isi",
		paymentMethodError:     "Metode pembayaran belum tersedia",
		courierUnavail:         "Kurir tidak bisa menjangkau alamat pengiriman",
		courierError:           "Kurir tidak tersedia",
		bodyPayload:            "Permintaan kamu tidak lengkap",
		paymentMethod:          "Metode pembayaran tidak ditemukan",
		paymentError:           "Gagal membuat pembayaran",
		courierNotFound:        "Kurir tidak ditemukan",
		userNotFound:           "User tidak ditemukan",
		userAddressNotFound:    "Alamat user tidak ditemukan",
		cartNotYours:           "Keranjang belanja bukan milik Anda",
	}

	systemCode = map[string]string{
		cartActive:             "44",
		paymentMethod:          "45",
		courierNotFound:        "46",
		userNotFound:           "47",
		userAddressNotFound:    "48",
		cartNotYours:           "49",
		getCheckoutSuccess:     "50",
		updateCheckoutSuccess:  "51",
		getCheckoutFailed:      "52",
		updateCheckoutFailed:   "53",
		deliveryAddressNotHave: "54",
		paymentMethodError:     "55",
		courierUnavail:         "56",
		courierUnavail:         "57",
		paymentError:           "58",
		finishCheckoutSuccess:  "60",
		finishCheckoutFailed:   "61",
		bodyPayload:            "80",
		queryFind:              "81",
		queryInsert:            "82",
		queryUpdate:            "83",
		keyRequired:            "84",
	}

	statusCode = map[string]int{
		keyRequired:            http.StatusBadRequest,
		getCheckoutSuccess:     http.StatusOK,
		getCheckoutFailed:      http.StatusBadRequest,
		updateCheckoutSuccess:  http.StatusOK,
		updateCheckoutFailed:   http.StatusBadRequest,
		queryFind:              http.StatusInternalServerError,
		queryInsert:            http.StatusInternalServerError,
		queryUpdate:            http.StatusInternalServerError,
		keyRequired:            http.StatusBadRequest,
		cartActive:             http.StatusInternalServerError,
		deliveryAddressNotHave: http.StatusNotFound,
		paymentMethodError:     http.StatusNotFound,
		courierUnavail:         http.StatusNotFound,
		courierError:           http.StatusBadRequest,
		bodyPayload:            http.StatusBadRequest,
		paymentMethod:          http.StatusNotFound,
		paymentError:           http.StatusInternalServerError,
		courierNotFound:        http.StatusNotFound,
		userNotFound:           http.StatusNotFound,
		userAddressNotFound:    http.StatusNotFound,
		cartNotYours:           http.StatusForbidden,
		finishCheckoutSuccess:  http.StatusCreated,
	}
)

func ResponseCheckout(isCheckoutProgress bool, value interface{}, err error) utils.ResponseInterface {
	if err != nil {
		var errMessage = err.Error()

		return utils.Error(statusCode[errMessage], systemCode[errMessage], message[errMessage], err)
	}

	cart, ok := value.(*model.Cart)
	if ok {
		checkout := createCheckoutView(cart)

		if isCheckoutProgress {
			return utils.Success(statusCode[updateCheckoutSuccess], systemCode[updateCheckoutSuccess], checkout)
		}

		return utils.Success(statusCode[getCheckoutSuccess], systemCode[getCheckoutSuccess], checkout)
	}

	if isCheckoutProgress {
		return utils.Success(statusCode[updateCheckoutSuccess], systemCode[updateCheckoutSuccess], value)
	}

	return utils.Success(statusCode[getCheckoutSuccess], systemCode[getCheckoutSuccess], value)
}

func createCheckoutView(cart *model.Cart) *checkoutView {
	var (
		wg                     sync.WaitGroup
		availableCouriers      []*availableCourierView
		availablePaymentMethod []*availablePaymentMethodView
		customerView           *customerView
		customerAddressView    *deliveryAddressView
		products               []*cartProductView
		courierView            *courierView
		paymentMethodView      *paymentMethodView
		deliveryCost           float64
		deliveryCostDiscount   float64
		totalProductsPrice     float64
		totalDeliveryCost      float64
		totalPayment           float64
	)

	wg.Add(1)

	go func() {
		defer wg.Done()
		availableCouriers = createAvailableCouriers(cart.GetAvailableCourier())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		customerView = createUserView(cart.GetUser())
		customerAddressView = createUserDeliveryAddress(cart.GetUserAddress())
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		products, totalProductsPrice = createCartProductView(cart)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		availablePaymentMethod = createAvailablePaymentMethods(cart.GetAvailablePaymentMethod())
	}()

	if cart.GetCourier() != nil || (cart.GetCourier() != nil && cart.GetCourier().Id != 0) {
		courierView = createCourierView(cart.GetCourier())

		totalDeliveryCost = courierView.TotalPrice
		deliveryCost = courierView.Price
		deliveryCostDiscount = courierView.DiscountedPrice
	}

	if cart.GetPaymentMethod() != nil || (cart.GetPaymentMethod() != nil && cart.GetPaymentMethod().Id != 0) {
		paymentMethodView = createPaymentMethod(cart.GetPaymentMethod())
	}

	wg.Wait()
	totalPayment = totalProductsPrice + totalDeliveryCost

	checkoutView := &checkoutView{
		Id:                      cart.Id,
		Key:                     cart.Key,
		TotalProductsPrice:      totalProductsPrice,
		DeliveryCost:            deliveryCost,
		DeliveryCostDiscount:    deliveryCostDiscount,
		TotalDeliveryCost:       totalDeliveryCost,
		TotalPayment:            totalPayment,
		Customer:                customerView,
		DeliveryAddress:         customerAddressView,
		Courier:                 courierView,
		PaymentMethod:           paymentMethodView,
		Products:                products,
		AvailableCouriers:       availableCouriers,
		AvailablePaymentMethods: availablePaymentMethod,
	}

	return checkoutView
}

func createUserView(user *model.User) *customerView {
	return &customerView{
		Id:          user.Id,
		Name:        user.Name,
		PhoneNumber: user.Phone,
	}
}

func createUserDeliveryAddress(deliveryAddress *model.UserAddress) *deliveryAddressView {
	deliveryAddressView := &deliveryAddressView{
		Id:           deliveryAddress.Id,
		ReceiverName: deliveryAddress.ReceiverName,
		PhoneNumber:  deliveryAddress.PhoneNumber,
		Address:      deliveryAddress.Address,
		PostCode:     deliveryAddress.PostCode,
		Lat:          deliveryAddress.Lat,
		Long:         deliveryAddress.Long,
		Province:     createRegionView(deliveryAddress.Province),
		City:         createRegionView(deliveryAddress.City),
		District:     createRegionView(deliveryAddress.District),
		Area:         createRegionView(deliveryAddress.Area),
	}

	return deliveryAddressView
}

func createRegionView(region *model.Region) *regionView {
	return &regionView{
		Id:   region.Id,
		Name: region.Name,
		Lat:  region.Lat,
		Long: region.Long,
	}
}
