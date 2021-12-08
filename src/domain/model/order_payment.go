package model

import "time"

type StatusPayment string

const (
	StatusPaymentWaiting  StatusPayment = "waiting_payment"
	StatusPaymentSuccess  StatusPayment = "success"
	StatusPaymentExpired  StatusPayment = "expired"
	StatusPaymentCanceled StatusPayment = "canceled"
	StatusPaymentFailed   StatusPayment = "failed"
)

type OrderPayment struct {
	Id              int
	OrderId         int
	PaymentMethodId int
	TotalPayment    float64
	TotalPaid       float64
	Status          StatusPayment
	IsActive        bool
	ConfirmedAt     *time.Time
	CreatedAt       time.Time
	UpdatedAt       *time.Time
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
