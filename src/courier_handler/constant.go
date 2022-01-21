package courierhandler

import "errors"

type courier int

const (
	_ courier = iota
	Shipper
)

const (
	domestic = "domestic"
	postpay  = "postpay"
)

var (
	errCreateDelivery           = errors.New("courier.create_delivery.error")
	errParseTimeout             = errors.New("parse.duration.error")
	errCourierNotYetImplemented = errors.New("courier.not_yet_implemented")
)
