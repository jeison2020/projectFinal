package routes

import (
	. "../handlers"
)

func registerUsersRoutes() {

	router.HandleFunc("/api/v001/users", CreateUser).Methods("POST") // Create

}