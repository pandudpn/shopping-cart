package model

import (
	"fmt"
	"strings"
)

type RequestShipperPricing struct {
	Cod         bool                   `json:"cod"`
	Limit       int                    `json:"limit"`
	Page        int                    `json:"page"`
	Weight      float64                `json:"weight"`
	Height      float64                `json:"height"`
	Length      int                    `json:"length"`
	Width       int                    `json:"width"`
	ItemValue   float64                `json:"item_value"`
	ForOrder    bool                   `json:"for_order"`
	SortBy      []string               `json:"sort_by"`
	Destination ShipperLocationRequest `json:"destination"`
	Origin      ShipperLocationRequest `json:"origin"`
}

type ResponseShipperPricing struct {
	Metadata   ShipperMetaData    `json:"metadat"`
	Data       DataShipperPricing `json:"data"`
	Pagination ShipperPagination  `json:"pagination"`
}

type ShipperLocationRequest struct {
	AreaId   int    `json:"area_id"`
	Lat      string `json:"lat"`
	Lng      string `json:"lng"`
	SuburbId int    `json:"suburb_id"`
}

type ShipperLocationResponse struct {
	AreaId       int     `json:"area_id"`
	AreaName     string  `json:"area_name,omitempty"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
	SuburbId     int     `json:"suburb_id"`
	SuburbName   string  `json:"suburb_name,omitempty"`
	CityId       int     `json:"city_id,omitempty"`
	CityName     string  `json:"city_name,omitempty"`
	ProvinceId   int     `json:"province_id,omitempty"`
	ProvinceName string  `json:"province_name,omitempty"`
}

type DataShipperPricing struct {
	Origin      ShipperLocationResponse `json:"origin"`
	Destination ShipperLocationResponse `json:"destination"`
	Pricings    []ShipperPricing        `json:"pricings"`
}

type ShipperPricing struct {
	Weight           float64         `json:"weight"`
	Volume           float64         `json:"volume"`
	VolumeWeight     float64         `json:"volume_weight"`
	FinalWeight      float64         `json:"final_weight"`
	MinDay           int             `json:"min_day"`
	MaxDay           int             `json:"max_day"`
	UnitPrice        float64         `json:"unit_price"`
	TotalPrice       float64         `json:"total_price"`
	Discount         float64         `json:"discount"`
	DiscountValue    float64         `json:"discount_value"`
	DiscountedPrice  float64         `json:"discounted_price"`
	InsuranceFee     float64         `json:"insurance_fee"`
	MustUseInsurance bool            `json:"must_use_insurance"`
	LiabilityValue   float64         `json:"liability_value"`
	FinalPrice       float64         `json:"final_price"`
	Currency         string          `json:"currency"`
	InsuranceApplied bool            `json:"insurance_applied"`
	Rate             ShipperRate     `json:"rate"`
	Logistic         ShipperLogistic `json:"logistic"`
}

type ShipperRate struct {
	Id              int    `json:"id"`
	Name            string `json:"name"`
	Type            string `json:"type"`
	Description     string `json:"description"`
	FullDescription string `json:"full_description"`
	IsHubless       bool   `json:"is_hubless"`
}

type ShipperLogistic struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	LogoUrl     string `json:"logo_url"`
	Code        string `json:"code"`
	CompanyName string `json:"company_name"`
}

type ShipperMetaData struct {
	Path           string `json:"path,omitempty"`
	HttpStatusCode int    `json:"http_status_code"`
	HttpStatus     string `json:"http_status"`
	Timestamp      int64  `json:"timestamp"`
}

type ShipperPagination struct {
	CurrentPage     int      `json:"current_page"`
	CurrentElements int      `json:"current_elements"`
	TotalPages      int      `json:"total_pages"`
	TotalElements   int      `json:"total_elements"`
	SortBy          []string `json:"sort_by"`
}

func NewRequestShipperPricing() RequestShipperPricing {
	return RequestShipperPricing{
		Limit:    100,
		Page:     1,
		Cod:      false,
		ForOrder: true,
		Height:   10,
		Width:    10,
		Length:   10,
		SortBy:   []string{"final_price"},
		Origin: ShipperLocationRequest{
			AreaId:   4756,
			SuburbId: 489,
			Lat:      "-6.2409898",
			Lng:      "106.763273",
		},
	}
}

func (sp ShipperPricing) GetCode() string {
	var rateType string
	switch sp.Rate.Type {
	case "Instant":
		rateType = "instant"
	case "Express":
		rateType = "nd"
	case "Same Day":
		rateType = "sd"
	default:
		rateType = "reg"
	}

	return fmt.Sprintf("%s%s", strings.ToLower(sp.Logistic.Code), rateType)
}
