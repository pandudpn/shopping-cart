package model

import "time"

type StatusPayment string

const (
	StatusPaymentWaiting  StatusPayment = "waiting_payment"
	StatusPaymentSuccess  StatusPayment = "success"
	StatusPaymentExpired  StatusPayment = "expired"
	StatusPaymentCanceled StatusPayment = "canceled"
	StatusPaymentFailed   StatusPayment = "failed"
	StatusPaymentRefunded StatusPayment = "refund"
)

type OrderPayment struct {
	Id              int
	OrderId         int
	PaymentMethodId int
	TotalPayment    float64
	TotalPaid       float64
	Status          StatusPayment
	IsActive        bool
	QrLink          *string
	RedirectLink    *string
	DeepLink        *string
	VaNumber        *string
	ExpiredAt       time.Time
	ConfirmedAt     *time.Time
	CreatedAt       time.Time
	UpdatedAt       *time.Time

	// Relation table
	PaymentMethod *PaymentMethod
}

func NewOrderPayment() *OrderPayment {
	now := time.Now().UTC()

	return &OrderPayment{
		Status:    StatusPaymentWaiting,
		IsActive:  true,
		CreatedAt: now,
	}
}

func (op *OrderPayment) SetStatusToWaiting() {
	op.Status = StatusPaymentWaiting
}

func (op *OrderPayment) SetStatusToSuccess() {
	op.Status = StatusPaymentSuccess
}

func (op *OrderPayment) SetStatusToFailed() {
	op.Status = StatusPaymentFailed
}

func (op *OrderPayment) SetStatusToExpired() {
	op.Status = StatusPaymentExpired
}

func (op *OrderPayment) SetStatusToCanceled() {
	op.Status = StatusPaymentCanceled
}

func (op *OrderPayment) SetStatusToRefund() {
	op.Status = StatusPaymentRefunded
}

func (op *OrderPayment) SetPaymentMethod(paymentMethod *PaymentMethod) {
	op.PaymentMethod = paymentMethod
}

func (op *OrderPayment) GetPaymentMethod() *PaymentMethod {
	return op.PaymentMethod
}
