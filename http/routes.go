/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package http

import (
	"github.com/gorilla/mux"
	v1 "test/http/rest/v1"
	"test/services"
	"test/utils"
)

func InitializeRoutes(con *utils.Helper) *mux.Router {
	r := mux.NewRouter()

	//Register Services
	var userService = services.NewUserService(con)

	//Register Controller
	userController := v1.NewUserController(userService)

	//Register Endpoints
	r.HandleFunc("/paygoal/users", userController.GetAllUsers).Methods("GET")
	return r
}
