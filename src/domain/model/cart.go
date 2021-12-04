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
	Products      []*CartProduct
	User          *User
	UserAddress   *UserAddress
	Courier       *Courier
	PaymentMethod *PaymentMethod

	// temporary data for cart and checkout
	Total                  float64
	TotalProductsPrice     float64
	TotalDeliveryCost      float64
	DeliveryCost           float64
	DeliveryCostDiscount   float64
	AvailableCourier       map[string]interface{}
	AvailablePaymentMethod []*PaymentMethod
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

func (c *Cart) SetUserAddress(userAddress *UserAddress) {
	c.UserAddress = userAddress
}

func (c *Cart) GetUserAddress() *UserAddress {
	return c.UserAddress
}

func (c *Cart) SetCourier(courier *Courier) {
	c.Courier = courier
}

func (c *Cart) GetCourier() *Courier {
	return c.Courier
}

func (c *Cart) SetPaymentMethod(paymentMethod *PaymentMethod) {
	c.PaymentMethod = paymentMethod
}

func (c *Cart) GetPaymentMethod() *PaymentMethod {
	return c.PaymentMethod
}

func (c *Cart) SetAvailableCourier(availableCourier map[string]interface{}) {
	c.AvailableCourier = availableCourier
}

func (c *Cart) GetAvailableCourier() map[string]interface{} {
	return c.AvailableCourier
}

func (c *Cart) SetAvailablePaymentMethod(availablePaymentMethod []*PaymentMethod) {
	c.AvailablePaymentMethod = availablePaymentMethod
}

func (c *Cart) GetAvailablePaymentMethod() []*PaymentMethod {
	return c.AvailablePaymentMethod
}

func (c *Cart) GetWeight() float64 {
	var weight float64

	for _, cartProduct := range c.GetProducts() {
		weight += cartProduct.GetProduct().Weight
	}

	weight = weight / 1000

	return weight
}
