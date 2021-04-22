package repository

import (
	. "../config"
	"../models"
)


func GetUserByUsername(username string) (user models.User) {
	db := GetConnection()
	db.Find(&user, "user_name = ?", username)
	return
}



func CreateUser(body models.User) (user models.User) {
	db := GetConnection()
	db.Create(&body)
	db.Last(&user)
	return
}
