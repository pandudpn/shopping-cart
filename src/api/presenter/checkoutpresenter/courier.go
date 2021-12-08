package checkoutpresenter

import (
	"sort"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/formatted"
)

func categoryCouriers() map[int]map[string]interface{} {
	return map[int]map[string]interface{}{
		0: {
			"category": model.Instant,
			"couriers": make([]*courierView, 0),
		},
		1: {
			"category": model.Sameday,
			"couriers": make([]*courierView, 0),
		},
		2: {
			"category": model.Nextday,
			"couriers": make([]*courierView, 0),
		},
		3: {
			"category": model.Regular,
			"couriers": make([]*courierView, 0),
		},
	}
}

func createAvailableCouriers(availableCouriers map[string]interface{}) []*availableCourierView {
	categoryCouriers := categoryCouriers()
	couriers := make([]*availableCourierView, 0)

	keyCategoryCouriers := make([]int, len(categoryCouriers))
	i := 0
	for key := range categoryCouriers {
		keyCategoryCouriers[i] = key
		i++
	}
	sort.Ints(keyCategoryCouriers)

	for key := range keyCategoryCouriers {
		couriersView := make([]*courierView, 0)
		var category string
		for _, availableCourier := range availableCouriers {
			if courier, ok := availableCourier.(*model.Courier); ok {
				if categoryCouriers[key]["category"] == courier.Category {
					category = courier.GetCategory()
					courierView := createCourierView(courier)

					couriersView = append(couriersView, courierView)
				}
			}
		}

		availCourier := &availableCourierView{
			Category: category,
			Couriers: couriersView,
		}

		couriers = append(couriers, availCourier)
	}

	return couriers
}

func createCourierView(courier *model.Courier) *courierView {
	courierView := &courierView{
		Id:              courier.Id,
		Code:            courier.Code,
		Name:            courier.Name,
		Category:        courier.Category,
		Price:           courier.DeliveryCost,
		DiscountedPrice: courier.DeliveryCostDiscount,
		InsurancePrice:  courier.DeliveryInsuranceCost,
		TotalPrice:      courier.TotalDeliveryCost,
		Image:           courier.Image,
		IsEligible:      courier.Enabled,
		Label:           courier.GetLabel(),
		Formatted: &formattedCourierView{
			Price:           formatted.IndonesiaCurrrency(courier.DeliveryCost),
			DiscountedPrice: formatted.IndonesiaCurrrency(courier.DeliveryCostDiscount),
			InsurancePrice:  formatted.IndonesiaCurrrency(courier.DeliveryInsuranceCost),
			TotalPrice:      formatted.IndonesiaCurrrency(courier.TotalDeliveryCost),
			Category:        courier.GetCategory(),
		},
	}

	return courierView
}
