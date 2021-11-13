package productimage

import "github.com/pandudpn/shopping-cart/src/domain/model"

const (
	QUERY_SELECT = "select pi.id, mf.filename, mf.url " +
		"from product_image as pi inner join media_file as mf on mf.id = pi.image_id "
	QUERY_PRODUCT_IMAGE = QUERY_SELECT + "where pi.product_id = $1 and mf.deleted = false and pi.deleted = false"
)

func (pi *ProductImageRepository) FindImagesByProductId(productId int) ([]*model.ProductImage, error) {
	var images = make([]*model.ProductImage, 0)

	rows, err := pi.DB.Query(QUERY_PRODUCT_IMAGE, productId)
	if err != nil {
		return images, err
	}
	defer rows.Close()

	for rows.Next() {
		image, err := rowsToImage(rows)
		if err != nil {
			return images, err
		}

		images = append(images, image)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return images, nil
}
