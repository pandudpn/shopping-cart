package mediafile

import (
	"database/sql"

	"github.com/pandudpn/shopping-cart/app/adapter/dbc"
	"github.com/pandudpn/shopping-cart/src/domain/model"
)

type MediaFileRepository struct {
	DB dbc.SqlDbc
}

func rowToFile(row *sql.Row) (*model.MediaFile, error) {
	mediaFile := &model.MediaFile{}

	err := row.Scan(&mediaFile.Id, &mediaFile.Filename, &mediaFile.Url)
	if err != nil {
		return nil, err
	}

	return mediaFile, nil
}
