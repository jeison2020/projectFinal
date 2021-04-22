package repository

import (
	. "../config"
	"../models"
)

func GetProfileModuleByIdProfiles(idProfile uint) (profile_module models.ProfileModule){
	db := GetConnection()
	db.Preload("Profile").Preload("Module").Find(&profile_module, "profile_id = ?", idProfile)
	return
}
