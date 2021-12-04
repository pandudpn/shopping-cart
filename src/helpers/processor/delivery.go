package processor

import (
	"encoding/json"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
)

func (p *processor) GetAvailableCourier(cart *model.Cart) error {
	var (
		couriers        = make(map[string]interface{})
		shipperCouriers model.ResponseShipperPricing
	)

	availableCouriers := p.getCouriers()
	if len(availableCouriers) < 1 {
		return ErrCourierNotAvail
	}

	couriersByte, err := p.courierRepo.GetCourierShipper(p.client, cart)
	if err != nil {
		logger.Log.Errorf("error get courier from shipper %v", err)
		return ErrCourier
	}
	err = json.Unmarshal(couriersByte, &shipperCouriers)
	if err != nil {
		logger.Log.Errorf("error unmarshal courier shipper %v", err)
		return ErrCourier
	}

	for _, shipment := range shipperCouriers.Data.Pricings {
		if cart.GetCourier() != nil && cart.GetCourier().Id != 0 {
			if cart.GetCourier().Code == shipment.GetCode() {
				cart.GetCourier().DeliveryCost = shipment.TotalPrice
				cart.GetCourier().DeliveryCostDiscount = shipment.DiscountedPrice
				cart.GetCourier().DeliveryInsuranceCost = shipment.InsuranceFee
				cart.GetCourier().TotalDeliveryCost = shipment.FinalPrice
				cart.GetCourier().MinDay = shipment.MinDay
				cart.GetCourier().MaxDay = shipment.MaxDay
				cart.GetCourier().Rate = shipment.Rate.Id

				cart.TotalDeliveryCost = shipment.FinalPrice
			}
		}

		for _, courier := range availableCouriers {
			if courier.Code == shipment.GetCode() {
				courier.DeliveryCost = shipment.TotalPrice
				courier.DeliveryCostDiscount = shipment.DiscountedPrice
				courier.DeliveryInsuranceCost = shipment.InsuranceFee
				courier.TotalDeliveryCost = shipment.FinalPrice
				courier.MinDay = shipment.MinDay
				courier.MaxDay = shipment.MaxDay
				courier.Rate = shipment.Rate.Id

				couriers[courier.Code] = courier

				break
			}
		}
	}

	cart.SetAvailableCourier(couriers)
	return nil
}

func (p *processor) getCouriers() []*model.Courier {
	couriers, err := p.courierRepo.FindEnabledCourier()
	if err != nil {
		logger.Log.Errorf("error get all couriers from database %v", err)
		return nil
	}

	for idx, courier := range couriers {
		if !courier.Enabled {
			couriers = append(couriers[:idx], couriers[idx+1:]...)
		}
	}

	return couriers
}
