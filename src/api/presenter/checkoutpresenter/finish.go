package checkoutpresenter

import (
	"sync"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils"
	"github.com/pandudpn/shopping-cart/src/utils/formatted"
)

type orderView struct {
	Id                 int                 `json:"id"`
	OrderNumber        string              `json:"orderNumber"`
	Status             string              `json:"status"`
	TotalProductsPrice float64             `json:"totalProductsPrice"`
	TotalDeliveryCost  float64             `json:"totalDeliveryCost"`
	TotalPayment       float64             `json:"totalPayment"`
	CreatedAt          string              `json:"createdAt"`
	ExpiredAt          string              `json:"expiredAt"`
	Formatted          *formattedOrderView `json:"formatted"`
	Payment            *paymentOrderView   `json:"payment"`
	Delivery           *deliveryOrderView  `json:"delivery"`
	Products           []*productOrderView `json:"products"`
}

type formattedOrderView struct {
	TotalProductsPrice string `json:"totalProductsPrice"`
	TotalDeliveryCost  string `json:"totalDeliveryCost"`
	TotalPayment       string `json:"totalPayment"`
	Status             string `json:"status"`
	CreatedAt          int64  `json:"createdAt"`
	ExpiredAt          int64  `json:"expiredAt"`
}

type paymentOrderView struct {
	Id            int                        `json:"id"`
	TotalPayment  float64                    `json:"totalPayment"`
	TotalPaid     float64                    `json:"totalPaid"`
	QrLink        *string                    `json:"qrLink"`
	RedirectLink  *string                    `json:"redirectLink"`
	DeepLink      *string                    `json:"deepLink"`
	Formatted     *formattedPaymentOrderView `json:"formatted"`
	PaymentMethod *paymentMethodView         `json:"paymentMethod"`
}

type formattedPaymentOrderView struct {
	TotalPayment string `json:"totalPayment"`
	TotalPaid    string `json:"totalPaid"`
}

type deliveryOrderView struct {
	Id                   int                         `json:"id"`
	DeliveryCost         float64                     `json:"deliveryCost"`
	DeliveryCostDiscount float64                     `json:"deliveryCostDiscount"`
	TotalDeliveryCost    float64                     `json:"totalDeliveryCost"`
	TrackingNumber       *string                     `json:"trackingNumber"`
	Status               string                      `json:"status"`
	Formatted            *formattedDeliveryOrderView `json:"formatted"`
	Courier              *courierView                `json:"courier"`
	UserAddress          *deliveryAddressView        `json:"customerAddress"`
}

type formattedDeliveryOrderView struct {
	DeliveryCost         string `json:"deliveryCost"`
	DeliveryCostDiscount string `json:"deliveryCostDiscount"`
	TotalDeliveryCost    string `json:"totalDeliveryCost"`
	Status               string `json:"status"`
}

type productOrderView struct {
	Id         int                       `json:"id"`
	Quantity   int                       `json:"quantity"`
	Price      float64                   `json:"price"`
	TotalPrice float64                   `json:"totalPrice"`
	Formatted  *formattedCartProductView `json:"formatted"`
	Product    *productView              `json:"product"`
}

func FinishCheckout(value interface{}, err error) utils.ResponseInterface {
	if err != nil {
		var errMsg = err.Error()
		return utils.Error(statusCode[errMsg], systemCode[finishCheckoutFailed], message[errMsg], err)
	}

	order := value.(*model.Order)

	var (
		wg           sync.WaitGroup
		paymentView  *paymentOrderView
		deliveryView *deliveryOrderView
		productsView []*productOrderView
	)

	wg.Add(1)
	go func() {
		defer wg.Done()
		paymentView = createOrderPaymentView(order.GetPayment())
		deliveryView = createDeliveryOrderView(order.GetDelivery())
		productsView = createProductOrderView(order.GetProducts())
	}()
	wg.Wait()

	status := string(order.Status)
	res := &orderView{
		Id:                 order.Id,
		OrderNumber:        order.OrderNumber,
		Status:             status,
		TotalProductsPrice: order.TotalProductsPrice,
		TotalDeliveryCost:  order.TotalDeliveryCost,
		TotalPayment:       order.TotalPayment,
		CreatedAt:          order.CreatedAt.In(timezone).Format(datetimeFormat),
		ExpiredAt:          order.GetPayment().ExpiredAt.In(timezone).Format(datetimeFormat),
		Formatted: &formattedOrderView{
			TotalProductsPrice: formatted.IndonesiaCurrrency(order.TotalProductsPrice),
			TotalDeliveryCost:  formatted.IndonesiaCurrrency(order.TotalDeliveryCost),
			TotalPayment:       formatted.IndonesiaCurrrency(order.TotalPayment),
			Status:             setStatus(status),
			CreatedAt:          order.CreatedAt.In(timezone).Unix(),
			ExpiredAt:          order.GetPayment().ExpiredAt.In(timezone).Unix(),
		},
		Payment:  paymentView,
		Delivery: deliveryView,
		Products: productsView,
	}

	return utils.Success(statusCode[finishCheckoutSuccess], systemCode[finishCheckoutSuccess], res)
}

func setStatus(status string) string {
	switch status {
	case "pending":
		return "Pending"
	case "waiting_payment":
		return "Menunggu Pembayaran"
	case "processing":
		return "Barang sedang diproses"
	case "delivered":
		return "Paket telah dikirim"
	case "package_received":
		return "Paket telah diterima"
	case "completed":
		return "Pesanan telah selesai"
	case "canceled":
		return "Pesanan dibatalkan"
	case "success", "payment_confirmed":
		return "Pembayaran berhasil"
	case "expired":
		return "Pembayaran telah kadaluarsa"
	case "failed":
		return "Pembayaran gagal"
	default:
		return ""
	}
}

func createOrderPaymentView(orderPayment *model.OrderPayment) *paymentOrderView {
	return &paymentOrderView{
		Id:           orderPayment.Id,
		TotalPayment: orderPayment.TotalPayment,
		TotalPaid:    orderPayment.TotalPaid,
		QrLink:       orderPayment.QrLink,
		RedirectLink: orderPayment.RedirectLink,
		DeepLink:     orderPayment.DeepLink,
		Formatted: &formattedPaymentOrderView{
			TotalPayment: formatted.IndonesiaCurrrency(orderPayment.TotalPayment),
			TotalPaid:    formatted.IndonesiaCurrrency(orderPayment.TotalPaid),
		},
		PaymentMethod: createPaymentMethod(orderPayment.GetPaymentMethod()),
	}
}

func createDeliveryOrderView(orderDelivery *model.OrderDelivery) *deliveryOrderView {
	status := string(orderDelivery.Status)
	return &deliveryOrderView{
		Id:                   orderDelivery.Id,
		DeliveryCost:         orderDelivery.DeliveryCost,
		DeliveryCostDiscount: orderDelivery.DeliveryCostDiscount,
		TotalDeliveryCost:    orderDelivery.TotalDeliveryCost,
		TrackingNumber:       orderDelivery.TrackingNumber,
		Status:               status,
		Formatted: &formattedDeliveryOrderView{
			DeliveryCost:         formatted.IndonesiaCurrrency(orderDelivery.DeliveryCost),
			DeliveryCostDiscount: formatted.IndonesiaCurrrency(orderDelivery.DeliveryCostDiscount),
			TotalDeliveryCost:    formatted.IndonesiaCurrrency(orderDelivery.TotalDeliveryCost),
			Status:               setStatus(status),
		},
		Courier:     createCourierView(orderDelivery.GetCourier()),
		UserAddress: createUserDeliveryAddress(orderDelivery.GetUserDelivery()),
	}
}

func createProductOrderView(orderProducts []*model.OrderProduct) []*productOrderView {
	productsOrderView := make([]*productOrderView, 0)

	for _, orderProduct := range orderProducts {
		productOrderView := &productOrderView{
			Id:         orderProduct.Id,
			Quantity:   orderProduct.Quantity,
			Price:      orderProduct.BasePrice,
			TotalPrice: orderProduct.TotalPrice,
			Formatted: &formattedCartProductView{
				Price:      formatted.IndonesiaCurrrency(orderProduct.BasePrice),
				TotalPrice: formatted.IndonesiaCurrrency(orderProduct.TotalPrice),
			},
			Product: createProductView(orderProduct.GetProduct()),
		}

		productsOrderView = append(productsOrderView, productOrderView)
	}

	return productsOrderView
}
