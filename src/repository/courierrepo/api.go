package courierrepo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/pandudpn/shopping-cart/src/domain/model"
	"github.com/pandudpn/shopping-cart/src/utils/logger"
	"github.com/spf13/viper"
)

func (cr *CourierRepository) GetCourierShipper(client *http.Client, cart *model.Cart) ([]byte, error) {
	var (
		apiKey     = viper.GetString("shipper.api.key")
		baseUrl    = viper.GetString("shipper.api.base")
		pricingUrl = viper.GetString("shipper.api.pricing")
		reqUrl     = baseUrl + pricingUrl
		reqPricing = model.NewRequestShipperPricing()
		destLat    = *cart.GetUserAddress().GetArea().Lat
		destLng    = *cart.GetUserAddress().GetArea().Long
	)
	header := map[string]string{
		"X-Api-Key":    apiKey,
		"Content-Type": "application/json",
	}

	if cart.GetUserAddress().Lat != nil && cart.GetUserAddress().Long != nil {
		destLat = *cart.GetUserAddress().Lat
		destLng = *cart.GetUserAddress().Long
	}

	reqPricing.ItemValue = cart.TotalProductsPrice
	reqPricing.Weight = cart.GetWeight()
	reqPricing.Destination = model.ShipperLocationRequest{
		AreaId:   cart.GetUserAddress().GetArea().RefId,
		SuburbId: cart.GetUserAddress().GetDistrict().RefId,
		Lat:      fmt.Sprintf("%f", destLat),
		Lng:      fmt.Sprintf("%f", destLng),
	}

	payload, err := json.Marshal(reqPricing)
	if err != nil {
		return nil, err
	}

	logger.Log.Debugf("req pricing %s", string(payload))

	req, err := http.NewRequest(http.MethodPost, reqUrl, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}

	for key, val := range header {
		req.Header.Set(key, val)
	}

	resBody, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resBody.Body.Close()

	if resBody.StatusCode >= 500 {
		err = fmt.Errorf("server.internal.error")
		return nil, err
	}

	body, err := ioutil.ReadAll(resBody.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}
