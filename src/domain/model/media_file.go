package model

import (
	"fmt"

	"github.com/spf13/viper"
)

type MediaFile struct {
	Id       int
	Filename *string
	Url      *string
}

func (mf *MediaFile) GetFile() string {
	if mf.Filename == nil && mf.Url != nil {
		return *mf.Url
	}

	return fmt.Sprintf("%s/%s", viper.GetString("aws.s3.bucket"), *mf.Filename)
}
