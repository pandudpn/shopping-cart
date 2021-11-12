package model

import "time"

type Product struct {
	Id              int
	CategoryId      int
	Name            string
	Slug            string
	Description     *string
	Price           float64
	DiscountedPrice float64
	Qty             int
	Enabled         bool
	CreatedAt       time.Time
	UpdatedAt       time.Time
	Deleted         bool

	// untuk relasi, silakan tambahkan struct dibawah
	// dan jangan lupa di inject pada setiap query
	Category *ProductCategory
	Images   []*ProductImage
}

func NewProduct() *Product {
	now := time.Now().UTC()

	return &Product{
		CreatedAt: now,
	}
}

func (p *Product) IsProductAvailable() bool {
	return p.Enabled && p.Qty > 0
}

func (p *Product) SetCategory(category *ProductCategory) {
	p.Category = category
}

func (p *Product) GetCategory() *ProductCategory {
	return p.Category
}

func (p *Product) SetImages(images []*ProductImage) {
	p.Images = images
}

func (p *Product) AddImage(image *ProductImage) {
	p.Images = append(p.Images, image)
}

func (p *Product) RemoveImage(image *ProductImage) {
	for idx, img := range p.Images {
		if img.Id == image.Id {
			p.Images = append(p.Images[:idx], p.Images[idx+1:]...)
			break
		}
	}
}

func (p *Product) GetImages() []*ProductImage {
	return p.Images
}
