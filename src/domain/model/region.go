package model

const (
	province = "PROVINCE"
	city     = "CITY"
	district = "DISTRICT"
	area     = "AREA"
)

type Region struct {
	Id       int
	ParentId *int
	RefId    int
	Name     string
	Category string
	Lat      *float64
	Long     *float64

	Parent *Region
}

func (r *Region) SetParent(parent *Region) {
	r.Parent = parent
}

func (r *Region) GetParent() *Region {
	return r.Parent
}

func (r *Region) IsParent() bool {
	return r.ParentId == nil
}

func (r *Region) IsProvince() bool {
	return r.Category == province
}

func (r *Region) IsCity() bool {
	return r.Category == city
}

func (r *Region) IsDistrict() bool {
	return r.Category == district
}

func (r *Region) IsArea() bool {
	return r.Category == area
}
