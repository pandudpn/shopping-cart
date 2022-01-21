package model

import "time"

type UserAddress struct {
	Id           int
	UserId       int
	ProvinceId   int
	CityId       int
	DistrictId   int
	AreaId       int
	ReceiverName string
	PhoneNumber  string
	PostCode     string
	Address      string
	Type         string
	Default      bool
	Lat          *float64
	Long         *float64
	CreatedAt    time.Time
	UpdatedAt    *time.Time
	Deleted      bool

	User     *User
	Province *Region
	City     *Region
	District *Region
	Area     *Region
}

func (ua *UserAddress) SetProvince(province *Region) {
	ua.Province = province
}

func (ua *UserAddress) GetProvince() *Region {
	return ua.Province
}

func (ua *UserAddress) SetCity(city *Region) {
	ua.City = city
}

func (ua *UserAddress) GetCity() *Region {
	return ua.City
}

func (ua *UserAddress) SetDistrict(district *Region) {
	ua.District = district
}

func (ua *UserAddress) GetDistrict() *Region {
	return ua.District
}

func (ua *UserAddress) SetArea(area *Region) {
	ua.Area = area
}

func (ua *UserAddress) GetArea() *Region {
	return ua.Area
}

func (ua *UserAddress) SetUser(user *User) {
	ua.User = user
}

func (ua *UserAddress) GetUser() *User {
	return ua.User
}

func (ua *UserAddress) GetLat() float64 {
	if ua.Lat == nil {
		return *ua.GetArea().Lat
	}

	return *ua.Lat
}

func (ua *UserAddress) GetLong() float64 {
	if ua.Long == nil {
		return *ua.GetArea().Long
	}

	return *ua.Long
}
