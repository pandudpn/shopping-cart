package model

import (
	"fmt"
	"time"
)

const (
	Sameday = "same_day"
	Instant = "instant"
	Regular = "regular"
	Nextday = "next_day"
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
	return c.Category == Sameday
}

func (c *Courier) IsInstant() bool {
	return c.Category == Instant
}

func (c *Courier) CanPickSameDayOrInstant(userAddress *UserAddress) bool {
	return userAddress.Lat != nil && userAddress.Long != nil
}

func (c *Courier) GetCategory() string {
	switch c.Category {
	case Instant:
		return "Instant"
	case Sameday:
		return "Same Day"
	case Nextday:
		return "Next Day"
	default:
		return "Regular"
	}
}

func (c *Courier) GetLabel() string {
	if c.IsInstant() || c.IsSameDay() {
		return "Estimasi tiba hari ini"
	} else {
		if c.MinDay == 1 && c.MaxDay == 1 {
			return "Estimasi tiba hari ini"
		} else if c.MinDay == 1 && c.MaxDay == 2 {
			return "Estimasi tiba besok"
		} else if c.MinDay == 1 && c.MaxDay > 2 {
			return fmt.Sprintf("Estimasi tiba besok - %s", addDate(c.MaxDay))
		} else if c.MinDay > 1 && c.MaxDay > 2 {
			return fmt.Sprintf("Estimasi tiba %s - %s", addDate(c.MinDay), addDate(c.MaxDay))
		}
	}

	return ""
}

func addDate(add int) string {
	timezone, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(timezone)

	return time.Date(now.Year(), now.Month(), now.Day()+add, 0, 0, 0, 0, timezone).Format("02 Jan")
}
