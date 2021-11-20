package model

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type Cart struct {
	Id              int
	UserId          int
	UserAddressId   *int
	CourierId       *int
	PaymentMethodId *int
	IsActive        bool
	Key             string
	CreatedAt       time.Time
	UpdatedAt       *time.Time

	// Relation table will be here
	Products []*CartProduct
	User     *User
}

func NewCart() *Cart {
	now := time.Now().UTC()

	key := uuid.New().String()
	key = strings.ReplaceAll(key, "-", "")
	return &Cart{
		IsActive:  true,
		Key:       key,
		CreatedAt: now,
	}
}

func (c *Cart) SetProducts(products []*CartProduct) {
	c.Products = products
}

func (c *Cart) AddProduct(product *CartProduct) {
	c.Products = append(c.Products, product)
}

func (c *Cart) RemoveProduct(product *CartProduct) {
	for idx, cartProduct := range c.Products {
		if product.Id == cartProduct.Id {
			c.Products = append(c.Products[:idx], c.Products[idx+1:]...)
			break
		}
	}
}

func (c *Cart) GetProducts() []*CartProduct {
	return c.Products
}

func (c *Cart) SetUser(user *User) {
	c.User = user
}

func (c *Cart) GetUser() *User {
	return c.User
}
