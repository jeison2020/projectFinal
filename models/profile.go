package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	ProfileId uint `json:"id"`
	ProName string `json:"name"`


}
