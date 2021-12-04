package model

import "time"

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
