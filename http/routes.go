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
	var groupService = services.NewGroupService(con.GetMariaDb())

	//Register Controller
	var userController = v1.NewUserController(userService)
	var groupController = v1.NewGroupController(groupService)

	//Register Endpoints
	r.Use(v1.CorsMiddleware)
	r.HandleFunc("/paygoal/users", userController.GetUsers).Methods("GET")
	r.HandleFunc("/paygoal/users", userController.CreateUser).Methods("POST")
	r.HandleFunc("/paygoal/users/", userController.GetUserByID).Methods("GET")
	r.HandleFunc("/paygoal/users/", userController.UpdateUserByID).Methods("PUT")
	r.HandleFunc("/paygoal/users/", userController.DeleteUserByID).Methods("DELETE")

	r.HandleFunc("/paygoal/groups", groupController.GetGroups).Methods("GET")
	r.HandleFunc("/paygoal/groups", groupController.CreateGroup).Methods("POST")
	return r
}
