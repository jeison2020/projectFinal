package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	UserId uint `json:"user_id"`
	UserName string `json:"user_name"`
	Hash string `json:"hash"
`

	
}
