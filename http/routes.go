/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package http

import (
	"github.com/gorilla/mux"
	"test/database"
	v1 "test/http/rest/v1"
	"test/services"
)

func InitializeRoutes(con *database.DB) *mux.Router {
	r := mux.NewRouter()

	//Register Services
	var userService = services.NewUserService(con.GetMariaDb())

	//Register Controller
	var userController = v1.NewUserController(userService)

	//Register Endpoints
	r.Use(v1.CorsMiddleware)
	r.HandleFunc("/paygoal/users", userController.GetAllUsers).Methods("GET")
	r.HandleFunc("/paygoal/users", userController.CreateUser).Methods("POST")
	r.HandleFunc("/paygoal/users/{id}", userController.GetUserByID).Methods("GET")

	return r
}
