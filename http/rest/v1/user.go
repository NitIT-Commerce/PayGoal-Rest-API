/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package v1

import (
	"encoding/json"
	"net/http"
	"test/models"
)

type UserController struct {
	User models.UserService
}

func NewUserController(user models.UserService) *UserController {
	return &UserController{
		User: user,
	}
}

func (user *UserController) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := user.User.GetAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	marshal, err := json.Marshal(users)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
	w.WriteHeader(http.StatusOK)
	return

}
