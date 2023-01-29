/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package v1

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"test/services"
)

type UserController struct {
	User *services.UserService
}

func NewUserController(user *services.UserService) *UserController {
	return &UserController{
		User: user,
	}
}

func (user *UserController) GetUserByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	users, err := user.User.GetUserByID(params["id"])
	log.Println(r.URL.Query().Get("id"))
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

func (user *UserController) CreateUser(w http.ResponseWriter, r *http.Request) {
	users, err := user.User.CreateUser(
		r.URL.Query().Get("user_email"),
		r.URL.Query().Get("user_pass"),
		r.URL.Query().Get("user_nickname"),
		r.URL.Query().Get("user_name"),
		r.URL.Query().Get("user_last_name"),
	)
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
