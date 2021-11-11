package model

type RequestRegister struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
}

type RequestLogin struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
