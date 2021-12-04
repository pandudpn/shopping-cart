package model

import "time"

const (
	sameday = "same_day"
	instant = "instant"
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
