package model

type ProductImage struct {
	Id        int
	ProductId int
	ImageId   int

	File *MediaFile
}

func (pi *ProductImage) SetImage(f *MediaFile) {
	pi.File = f
}

func (pi *ProductImage) GetImage() *MediaFile {
	return pi.File
}
