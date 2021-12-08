package model

import "time"

type StatusOrder string

const (
	StatusOrderWaitingPayment   StatusOrder = "waiting_payment"
	StatusOrderPaymentConfirmed StatusOrder = "payment_confirmed"
	StatusOrderProcessing       StatusOrder = "processing"
	StatusOrderDelivered        StatusOrder = "delivered"
	StatusOrderPackageReceived  StatusOrder = "package_received"
	StatusOrderCompleted        StatusOrder = "completed"
	StatusOrderReturn           StatusOrder = "return"
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
