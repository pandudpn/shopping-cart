package model

const (
	Ewallet    = "e-wallet"
	VA         = "virtual_account"
	CreditCard = "credit_card"

	MdGopay   = "midtransgopay"
	MdVa      = "midtransva"
	XdEwallet = "xenditewallet"
)

type PaymentMethod struct {
	Id       int
	Code     string
	Category string
	Name     string
	Image    *string
	Enabled  bool
}

func NewPaymentMethod() *PaymentMethod {
	return &PaymentMethod{
		Enabled: true,
	}
}

func (pm *PaymentMethod) GetCategory() string {
	switch pm.Category {
	case Ewallet:
		return "E-Wallet"
	case VA:
		return "Virtual Account"
	case CreditCard:
		return "Credit Card"
	default:
		return ""
	}
}
