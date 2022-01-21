package courierhandler

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
	"github.com/spf13/viper"
)

type shipper struct{}

func (s *shipper) Process(order *model.Order) error {
	baseUrl := viper.GetString("shipper.api.base")
	apiKey := viper.GetString("shipper.api.key")
	defaultAreaId := 4756
	defaultLat := -6.326193
	defaultLong := 106.892194
	timeout, err := time.ParseDuration(viper.GetString("application.timeout"))
	if err != nil {
		logger.Log.Errorf("error parse duration %v", err)
		return errParseTimeout
	}

	orderDelivery := order.GetDelivery()
	customer := order.GetUser()

	packages, totalPrice, totalWeight := s.products(order)
	logger.Log.Infof("total price = %f and total weight = %f", totalPrice, totalWeight)

	logger.Log.Debugf("rate of order delivery %d", *orderDelivery.Rate)
	reqShipper := &model.RequestShipperOrder{
		ExternalId:  fmt.Sprintf("%d", orderDelivery.Id),
		Coverage:    domestic,
		PaymentType: postpay,
		Courier: &model.ShipperCourier{
			RateId:       *orderDelivery.Rate,
			UseInsurance: true,
			Cod:          false,
		},
		Origin: &model.ShipperAddress{
			AreaId:  defaultAreaId,
			Lat:     fmt.Sprintf("%f", defaultLat),
			Lng:     fmt.Sprintf("%f", defaultLong),
			Address: "Jl. H. Mairin RT010 / RW003 No.2",
		},
		Destination: &model.ShipperAddress{
			AreaId:  orderDelivery.GetUserDelivery().GetArea().RefId,
			Address: orderDelivery.Address,
			Lat:     fmt.Sprintf("%f", orderDelivery.GetUserDelivery().GetLat()),
			Lng:     fmt.Sprintf("%f", orderDelivery.GetUserDelivery().GetLong()),
		},
		Sender: &model.ShipperAssign{
			Name:        "Pandu dwi Putra Nugroho",
			PhoneNumber: "082122274139",
		},
		Received: &model.ShipperAssign{
			Name:        customer.Name,
			PhoneNumber: customer.Phone,
		},
		Package: &model.ShipperPackage{
			Price:       totalPrice,
			Height:      10,
			Width:       10,
			Length:      10,
			Weight:      totalWeight,
			Items:       packages,
			PackageType: 2,
		},
	}

	buf, err := json.Marshal(reqShipper)
	if err != nil {
		return errCreateDelivery
	}

	logger.Log.Debugf("request new order %s", string(buf))

	url := fmt.Sprintf("%s%s", baseUrl, viper.GetString("shipper.api.order"))
	client := &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewReader(buf))
	if err != nil {
		logger.Log.Errorf("error request to shipper %v", err)
		return errCreateDelivery
	}
	req.Header.Set("x-api-key", apiKey)

	res, err := client.Do(req)
	if err != nil {
		logger.Log.Errorf("error request to shipper %v", err)
		return errCreateDelivery
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Log.Error(err)
		return errCreateDelivery
	}
	defer res.Body.Close()

	logger.Log.Infof("response from shipper %s", string(body))

	if res.StatusCode >= 400 {
		return errCreateDelivery
	}

	resShipper := model.ResponseShipperOrder{}
	err = json.Unmarshal(body, &resShipper)
	if err != nil {
		logger.Log.Error(err)
		return errCreateDelivery
	}

	orderDelivery.RefId = &resShipper.Data.OrderId

	return nil
}

func (s *shipper) products(order *model.Order) ([]*model.ShipperPackageItem, float64, float64) {
	var (
		totalPrice,
		totalWeight float64
		items = make([]*model.ShipperPackageItem, 0)
	)

	for _, product := range order.GetProducts() {
		totalPrice += product.TotalPrice
		totalWeight += product.GetProduct().Weight * float64(product.Quantity)

		item := &model.ShipperPackageItem{
			Name:  product.GetProduct().Name,
			Price: product.BasePrice,
			Qty:   product.Quantity,
		}

		items = append(items, item)
	}

	return items, totalPrice, totalWeight / 1000
}
