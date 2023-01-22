/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package http

import (
	"github.com/gorilla/mux"
	"log"
	"test/database"
	v1 "test/http/rest/v1"
	"test/services"
)

func InitializeRoutes(con *database.Container) *mux.Router {
	r := mux.NewRouter()

	log.Println("Router izz da")

	//Register Services
	var userService = services.NewUserService(con.GetSqlConnection())

	log.Println("Service izz da")

	//Register Controller
	userController := v1.NewUserController(userService)
	log.Println("Controller izz da")

	//Register Endpoints
	r.Use(v1.CorsMiddleware)

	r.HandleFunc("/paygoal/users", userController.GetAllUsers).Methods("GET")
	log.Println("Endpoint izz da")

	return r
}
