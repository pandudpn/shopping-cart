package model

import (
	"time"
)

type StatusOrder string

const (
	StatusOrderWaitingPayment   StatusOrder = "waiting_payment"
	StatusOrderPaymentConfirmed StatusOrder = "payment_confirmed"
	StatusOrderProcessing       StatusOrder = "processing"
	StatusOrderDelivered        StatusOrder = "delivered"
	StatusOrderPackageReceived  StatusOrder = "package_received"
	StatusOrderCompleted        StatusOrder = "completed"
	StatusOrderReturn           StatusOrder = "return"
	StatusOrderFailed           StatusOrder = "failed"
	StatusOrderExpired          StatusOrder = "expired"
)

type Order struct {
	Id                 int
	CartId             int
	UserId             int
	DeliveryAddressId  int
	OrderNumber        string
	TotalProductsPrice float64
	TotalDeliveryCost  float64
	TotalPayment       float64
	Status             StatusOrder
	CompletedAt        *time.Time
	CanceledAt         *time.Time
	CreatedAt          time.Time
	UpdatedAt          *time.Time

	// Relation table
	Products []*OrderProduct
	Payment  *OrderPayment
	Delivery *OrderDelivery
	User     *User
}

func NewOrder() *Order {
	now := time.Now().UTC()

	return &Order{
		Status:    StatusOrderWaitingPayment,
		CreatedAt: now,
	}
}

func (o *Order) SetStatusToWaitingPayment() {
	o.Status = StatusOrderWaitingPayment
}

func (o *Order) SetStatusToPaymentConfirmed() {
	o.Status = StatusOrderPaymentConfirmed
}

func (o *Order) SetStatusToProcessing() {
	o.Status = StatusOrderProcessing
}

func (o *Order) SetStatusToDelivered() {
	o.Status = StatusOrderDelivered
}

func (o *Order) SetStatusToPackageReceived() {
	o.Status = StatusOrderPackageReceived
}

func (o *Order) SetStatusToCompleted() {
	o.Status = StatusOrderCompleted
}

func (o *Order) SetStatusToReturn() {
	o.Status = StatusOrderReturn
}

func (o *Order) SetStatusToFailed() {
	o.Status = StatusOrderFailed
}

func (o *Order) SetStatusToExpired() {
	o.Status = StatusOrderExpired
}

func (o *Order) AddProduct(product *OrderProduct) {
	o.Products = append(o.Products, product)
}

func (o *Order) RemoveProduct(product *OrderProduct) {
	for idx, op := range o.Products {
		if op.Id == product.Id {
			o.Products = append(o.Products[:idx], o.Products[idx+1:]...)
			break
		}
	}
}

func (o *Order) SetProducts(products []*OrderProduct) {
	o.Products = products
}

func (o *Order) GetProducts() []*OrderProduct {
	return o.Products
}

func (o *Order) SetPayment(payment *OrderPayment) {
	o.Payment = payment
}

func (o *Order) GetPayment() *OrderPayment {
	return o.Payment
}

func (o *Order) SetDelivery(delivery *OrderDelivery) {
	o.Delivery = delivery
}

func (o *Order) GetDelivery() *OrderDelivery {
	return o.Delivery
}

func (o *Order) SetUser(user *User) {
	o.User = user
}

func (o *Order) GetUser() *User {
	return o.User
}
