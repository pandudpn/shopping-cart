package model

import "time"

type CartProduct struct {
	Id         int
	CartId     int
	ProductId  int
	Quantity   int
	BasePrice  float64
	TotalPrice float64
	CreatedAt  time.Time

	Product *Product
	Cart    *Cart
}

func NewCartProduct() *CartProduct {
	now := time.Now().UTC()
	return &CartProduct{
		CreatedAt: now,
	}
}

func (cp *CartProduct) SetCart(cart *Cart) {
	cp.Cart = cart
}

func (cp *CartProduct) GetCart() *Cart {
	return cp.Cart
}

func (cp *CartProduct) SetProduct(product *Product) {
	cp.Product = product
}

func (cp *CartProduct) GetProduct() *Product {
	return cp.Product
}
