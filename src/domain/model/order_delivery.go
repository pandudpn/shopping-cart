package model

import "time"

type StatusDelivery string

const (
	StatusDeliveryPending         StatusDelivery = "pending"
	StatusDeliveryProcessing      StatusDelivery = "processing"
	StatusDeliveryDelivered       StatusDelivery = "delivered"
	StatusDeliveryPackageReceived StatusDelivery = "package_received"
	StatusDeliveryCanceled        StatusDelivery = "canceled"
)

type OrderDelivery struct {
	Id                   int
	OrderId              int
	CourierId            int
	DeliveryAddressId    int
	DeliveryCost         float64
	DeliveryCostDiscount float64
	TotalDeliveryCost    float64
	TrackingNumber       *string
	Address              string
	ReceivedName         string
	PhoneNumber          string
	Lat                  *float64
	Long                 *float64
	Status               StatusDelivery
	DeliveredAt          *time.Time
	PackageReceivedAt    *time.Time
	CreatedAt            time.Time
	UpdatedAt            *time.Time
	Rate                 *int
	RefId                *string

	// Relation table
	Courier      *Courier
	UserDelivery *UserAddress
	Statuses     []*OrderDeliveryStatus
}

func NewOrderDelivery() *OrderDelivery {
	now := time.Now().UTC()

	return &OrderDelivery{
		Status:    StatusDeliveryPending,
		CreatedAt: now,
	}
}

func (od *OrderDelivery) SetStatusToPending() {
	od.Status = StatusDeliveryPending
}

func (od *OrderDelivery) SetStatusToDelivered() {
	od.Status = StatusDeliveryDelivered
}

func (od *OrderDelivery) SetStatusToPackageReceived() {
	od.Status = StatusDeliveryPackageReceived
}

func (od *OrderDelivery) SetStatusToCanceled() {
	od.Status = StatusDeliveryCanceled
}

func (od *OrderDelivery) SetStatusToProcessing() {
	od.Status = StatusDeliveryProcessing
}

func (od *OrderDelivery) SetCourier(courier *Courier) {
	od.Courier = courier
}

func (od *OrderDelivery) GetCourier() *Courier {
	return od.Courier
}

func (od *OrderDelivery) SetUserDelivery(userAddress *UserAddress) {
	od.UserDelivery = userAddress
}

func (od *OrderDelivery) GetUserDelivery() *UserAddress {
	return od.UserDelivery
}

func (od *OrderDelivery) AddStatus(status *OrderDeliveryStatus) {
	od.Statuses = append(od.Statuses, status)
}

func (od *OrderDelivery) RemoveStatus(status *OrderDeliveryStatus) {
	for idx, o := range od.Statuses {
		if o.Id == status.Id {
			od.Statuses = append(od.Statuses[:idx], od.Statuses[idx+1:]...)
			break
		}
	}
}

func (od *OrderDelivery) SetStatuses(statuses []*OrderDeliveryStatus) {
	od.Statuses = statuses
}

func (od *OrderDelivery) GetStatuses() []*OrderDeliveryStatus {
	return od.Statuses
}
