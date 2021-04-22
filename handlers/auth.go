package handlers

import (
	"net/http"

	"golang.org/x/crypto/bcrypt"

	"../io/request"
	"../io/response"
	"../jwt"
	"../models"
	rep "../repository"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var body models.User
	err := request.Json(r, &body)
	if err != nil {
		response.Error("Malformed body", http.StatusBadRequest, w)
		return
	}
	if len(body.UserName) <= 0 || len(body.Hash) <= 0 {
		response.Error("Email and password required", http.StatusBadRequest, w)
		return
	}
	user := rep.GetUserByUsername(body.UserName)
	if user.ID <= 0 {
		response.Error("User does not exist", http.StatusBadRequest, w)
		return
	}
	userPass := []byte(user.Hash)
	bodyPass := []byte(body.Hash)
	err = bcrypt.CompareHashAndPassword(userPass, bodyPass)
	if err != nil {
		response.Error("Wrong password", http.StatusBadRequest, w)
		return
	}
	token, err := jwt.GenerateToken(user)
	if err != nil {
		response.Error("Error generating JWT Token", http.StatusInternalServerError, w)
		return
	}
	user.Hash = ""
	loginResp := models.LoginResponse{Token: token, User: user}
	response.Json(loginResp, http.StatusOK, w)
}