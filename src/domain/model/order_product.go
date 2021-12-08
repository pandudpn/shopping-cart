package model

import "time"

type OrderProduct struct {
	Id         int
	OrderId    int
	ProductId  int
	Quantity   int
	BasePrice  float64
	TotalPrice float64
	CreatedAt  time.Time
}

func NewOrderProduct() *OrderProduct {
	now := time.Now().UTC()
	return &OrderProduct{
		CreatedAt: now,
	}
}
