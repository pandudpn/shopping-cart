package model

import "time"

const (
	sameday = "same_day"
	instant = "instant"
	regular = "regular"
	nextday = "next_day"
)

type Courier struct {
	Id        int
	Code      string
	Name      string
	Category  string
	Image     *string
	Enabled   bool
	CreatedAt time.Time
	UpdatedAt *time.Time
	Deleted   bool

	// temp variable
	Rate                  int
	MinDay                int
	MaxDay                int
	DeliveryCost          float64 // base price dari delivery cost courier
	DeliveryCostDiscount  float64
	DeliveryInsuranceCost float64
	TotalDeliveryCost     float64 // total delivery adalah penjumlahan ataupun pengurangan dari (delivery_cost + delivery_insurance_cost - delivery_cost_discount)
}

func NewCourier() *Courier {
	now := time.Now().UTC()
	return &Courier{
		Enabled:   true,
		CreatedAt: now,
	}
}

func (c *Courier) SetImage(name string) {
	c.Image = &name
}

func (c *Courier) GetImage() string {
	if c.Image == nil {
		return ""
	}
	return *c.Image
}

func (c *Courier) IsSameDay() bool {
	return c.Category == sameday
}

func (c *Courier) IsInstant() bool {
	return c.Category == instant
}

func (c *Courier) CanPickSameDayOrInstant(userAddress *UserAddress) bool {
	return userAddress.Lat != nil && userAddress.Long != nil
}

func (c *Courier) GetCategory() string {
	switch c.Category {
	case instant:
		return "Instant"
	case sameday:
		return "Same Day"
	case nextday:
		return "Next Day"
	default:
		return "Regular"
	}
}
