package models

import "gorm.io/gorm"

type Person struct {
	gorm.Model
	PersonId uint `json:"person_id"`
	PerName string `json:"name"`
	PerAge string `json:"age"`
	PerTypeDoc string `json:"type_doc"`
	PerNumberDoc string `json:"number_doc"`


}
