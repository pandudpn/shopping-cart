package productimage

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type ProductImageRepository struct {
	DB dbc.SqlDbc
}

func rowsToImage(rows *sql.Rows) (*model.ProductImage, error) {
	image := &model.ProductImage{}
	mediaFile := &model.MediaFile{}

	err := rows.Scan(&image.Id, &mediaFile.Filename, &mediaFile.Url)
	if err != nil {
		return nil, err
	}

	image.SetImage(mediaFile)
	return image, nil
}
