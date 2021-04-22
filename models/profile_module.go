package models

import "gorm.io/gorm"

type ProfileModule struct {
	gorm.Model
	ProfileId  uint   `json:"profile_id"`
	Profile Profile `gorm:"foreignKey:ProfileId" json:"profile,omitempty"`
	ModuleId uint `json:"module_id"`
	Module Module `gorm:"foreignKey:ModuleId" json:"module,omitempty"`


}
