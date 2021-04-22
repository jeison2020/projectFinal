package routes

import (
	. "../handlers"
)

func registerAuthRoutes() {
	router.HandleFunc("/api/v001/auth", Login).Methods("POST")
}