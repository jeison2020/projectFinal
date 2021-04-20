package models

import "gorm.io/gorm"

type Module struct {
	gorm.Model
	ModuleId uint `json:"id"`
	ModName string `json:"name"`
}
