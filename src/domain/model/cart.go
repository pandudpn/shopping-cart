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
