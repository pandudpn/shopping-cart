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

type RequestAddToCart struct {
	ProductId int `json:"productId"`
	Quantity  int `json:"quantity"`
	SecretKey string
	UserId    int
}

type RequestCheckout struct {
	UserId          int
	CartKey         string
	Courier         *RequestUpdateCheckout `json:"courier"`
	PaymentMethod   *RequestUpdateCheckout `json:"paymentMethod"`
	DeliveryAddress *RequestUpdateCheckout `json:"deliveryAddress"`
}

type RequestUpdateCheckout struct {
	Id int `json:"id"`
}
