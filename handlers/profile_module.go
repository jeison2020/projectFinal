package handlers

import (
	"../io/response"
	rep "../repository"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

func GetProfileModuleByIdProfiles(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	if id, ok := vars["id"]; ok {
		userId, err := strconv.Atoi(id)
		if err != nil {
			response.Error("Given ID is not allowed", http.StatusBadRequest, w)
			return
		}
		profile_module := rep.GetProfileModuleByIdProfiles(uint(userId))
		status := http.StatusOK
		if profile_module.ID == 0 {
			status = http.StatusNoContent
		}
		response.Json(profile_module, status, w)
	} else {
		response.Error("", http.StatusNotFound, w)
	}

}
