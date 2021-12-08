package model

import "time"

type OrderDeliveryStatus struct {
	Id              int
	OrderDeliveryId int
	Status          StatusDelivery
	CreatedAt       time.Time
}

func NewOrderDeliveryStatus() *OrderDeliveryStatus {
	now := time.Now().UTC()
	return &OrderDeliveryStatus{
		CreatedAt: now,
	}
}
