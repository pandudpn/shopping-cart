package model

import "time"

type Stock struct {
	Id           int
	ProductId    int
	QuantityHold int
	CreatedAt    time.Time
	UpdatedAt    *time.Time

	// relation table
	Product *Product
}

func NewStock() *Stock {
	now := time.Now().UTC()
	return &Stock{
		CreatedAt: now,
	}
}

func (s *Stock) SetProduct(product *Product) {
	s.Product = product
}

func (s *Stock) GetProduct() *Product {
	return s.Product
}
