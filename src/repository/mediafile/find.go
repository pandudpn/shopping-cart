package mediafile

import "github.com/pandudpn/shopping-cart/src/domain/model"

const (
	QUERY_SELECT      = "select id, filename, url from media_file"
	QUERY_IMAGE_BY_ID = QUERY_SELECT + "where id = $1 and deleted = false"
)

func (mfr *MediaFileRepository) FindImageById(id int) (*model.MediaFile, error) {
	row := mfr.DB.QueryRow(QUERY_IMAGE_BY_ID, id)

	mediaFile, err := rowToFile(row)
	if err != nil {
		return nil, err
	}

	return mediaFile, nil
}
