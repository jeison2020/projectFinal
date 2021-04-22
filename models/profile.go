package models

import "gorm.io/gorm"

type Profile struct {
	gorm.Model
	ProName string `json:"name"`


}
