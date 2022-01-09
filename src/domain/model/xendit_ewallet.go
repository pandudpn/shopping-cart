package model

const (
	StatusSucceeded = "SUCCEEDED"
	StatusFailed    = "FAILED"
	StatusVoided    = "VOIDED"
	StatusRefunded  = "REFUNDED"
)

type EWalletPaymentNotification struct {
	Event      string                  `json:"event"`
	Created    string                  `json:"created"`
	BusinessId string                  `json:"business_id"`
	Data       DataEWalletNotification `json:"data"`
}

type DataEWalletNotification struct {
	Id             string  `json:"id"`
	ReferenceId    string  `json:"reference_id"`
	Status         string  `json:"status"`
	Currency       string  `json:"currency"`
	ChargeAmount   float64 `json:"charge_amount"`
	CaptureAmount  float64 `json:"capture_amount"`
	ChannelCode    string  `json:"channel_code"`
	CheckoutMethod string  `json:"checkout_method"`
	FailedCode     string  `json:"failed_code"`
}

func (d *DataEWalletNotification) IsStatusSuccess() bool {
	return d.Status == StatusSucceeded
}

func (d *DataEWalletNotification) IsStatusFailed() bool {
	return d.Status == StatusFailed
}

func (d *DataEWalletNotification) IsStatusVoided() bool {
	return d.Status == StatusVoided
}

func (d *DataEWalletNotification) IsStatusRefund() bool {
	return d.Status == StatusRefunded
}
