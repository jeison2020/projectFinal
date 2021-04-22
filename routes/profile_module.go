package routes

import (
	. "../handlers"
)

func registrarProfileModuleRoutes()  {

	router.HandleFunc("/api/v001/profilemodule/{id}", GetProfileModuleByIdProfiles).Methods("GET")

}
