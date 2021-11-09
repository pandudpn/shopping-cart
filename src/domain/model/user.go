package model

import (
	"time"
)

type User struct {
	Id              int
	Name            string
	Email           string
	Phone           string
	Password        string
	Enabled         bool
	EmailVerifiedAt *time.Time
	CreatedAt       time.Time
}

func NewUser() *User {
	now := time.Now().UTC() // harus selalu jadi utc. jika ingin convert ke utc+7, in condition needed, tambahkan saja

	return &User{
		CreatedAt: now,
	}
}

func (u *User) IsVerified() bool {
	return u.EmailVerifiedAt != nil
}
