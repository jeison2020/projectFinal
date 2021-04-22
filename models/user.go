package models

import (
	"gorm.io/gorm"
)


type Rol string

const (
	SUPER  Rol = "SUPER"
	ADMIN Rol = "ADMIN"
	USER Rol = "USER"

)



type User struct {
	gorm.Model
	UserName string `json:"user_name"`
	Hash string `json:"user_hash"`
	Rol Rol `sql:"type:rol" json:"rol"`
	ProfileId  uint   `json:"profile_id"`
	Profile   Profile `gorm:"foreignKey:ProfileId" json:"profile,omitempty"`
}

