package model

type RequestShipperOrder struct {
	ExternalId  string          `json:"external_id"`
	Coverage    string          `json:"coverage"`
	PaymentType string          `json:"payment_type"`
	Courier     *ShipperCourier `json:"courier"`
	Origin      *ShipperAddress `json:"origin"`
	Destination *ShipperAddress `json:"destination"`
	Sender      *ShipperAssign  `json:"consignee"`
	Received    *ShipperAssign  `json:"consigner"`
	Package     *ShipperPackage `json:"package"`
}

type ResponseShipperOrder struct {
	Data struct {
		OrderId string `json:"order_id"`
		RequestShipperOrder
	} `json:"data"`
}

type ShipperCourier struct {
	Cod             bool    `json:"cod"`
	RateId          int     `json:"rate_id"`
	UseInsurance    bool    `json:"use_insurance"`
	Amount          float64 `json:"amount,omitempty"`
	InsuranceAmount float64 `json:"insurance_amount"`
}

type ShipperAssign struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
}

type ShipperPackage struct {
	Height      float64               `json:"height,omitempty"`
	Length      float64               `json:"length,omitempty"`
	PackageType int                   `json:"package_type,omitempty"`
	Price       float64               `json:"price"`
	Weight      float64               `json:"weight"`
	Width       float64               `json:"width,omitempty"`
	Items       []*ShipperPackageItem `json:"items"`
}

type ShipperPackageItem struct {
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Qty   int     `json:"qty"`
}

type ShipperAddress struct {
	AreaId  int    `json:"area_id"`
	Lat     string `json:"lat"`
	Lng     string `json:"lng"`
	Address string `json:"address"`
}
