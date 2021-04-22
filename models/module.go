package models

import "gorm.io/gorm"

type Module struct {
	gorm.Model
	ModName string `json:"name"`
}
