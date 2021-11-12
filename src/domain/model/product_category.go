package model

import "time"

type ProductCategory struct {
	Id        int
	Name      string
	Slug      string
	Enabled   bool
	CreatedAt time.Time
}

func NewProductCategory() *ProductCategory {
	now := time.Now().UTC()

	return &ProductCategory{
		CreatedAt: now,
	}
}
