package controller

import (
	"encoding/json"
	"net/http"
	"strconv"
	"subproblem/rest-api/pkg/service"

	"github.com/gorilla/mux"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		service: userService,
	}
}


// /api/v1/user/{id}

func (c *UserController) GetUserById(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	userId, err := strconv.Atoi(vars["id"])
	
	if err != nil {
		http.Error(w, "Incorrect user ID", http.StatusBadRequest)
		return
	}

	user, err := c.service.GetUserById(userId)
	
	if err != nil {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	userJson, err := json.Marshal(user)

	if err != nil {
		http.Error(w, "Error serializing data", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(userJson)
	

}
