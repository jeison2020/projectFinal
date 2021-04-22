package handlers

import (
	"errors"
	"net/http"

	"../crypto"
	"../io/request"
	"../io/response"
	"../models"
	rep "../repository"
)


func CreateUser(w http.ResponseWriter, r *http.Request) {
	var body models.User
	request.Json(r, &body)
	err := validateUser(body)
	if err != nil {
		response.Error(err.Error(), http.StatusBadRequest, w)
		return
	}
	existingUser := rep.GetUserByUsername(body.UserName)
	if existingUser.ID > 0 {
		response.Error("Duplicated Username", http.StatusBadRequest, w)
		return
	}
	pass, err := crypto.EncryptText(body.Hash)
	if err != nil {
		response.Error("Error encrypting password", http.StatusBadRequest, w)
		return
	}
	body.Hash = pass
	user := rep.CreateUser(body)
	response.Json(user, http.StatusCreated, w)
}


func validateUser(user models.User) (err error) {
	if len(user.UserName) < 4 {
		err = errors.New("Name property must have 4 characters at least")
		return
	}
	if len(user.Hash) < 6 {
		err = errors.New("Password must have 6 characters at least")
		return
	}
	return
}

