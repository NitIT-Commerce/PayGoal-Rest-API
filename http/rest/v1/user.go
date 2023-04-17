/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package v1

import (
	"encoding/json"
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
	users, err := user.User.GetUserByID(r.URL.Query().Get("id"))

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

func (user *UserController) DeleteUserByID(w http.ResponseWriter, r *http.Request) {
	err := user.User.DeleteUserByID(r.URL.Query().Get("id"))

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

func (user *UserController) UpdateUserByID(w http.ResponseWriter, r *http.Request) {
	if r.URL.Query().Get("id") == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id := r.URL.Query().Get("id")

	if r.URL.Query().Get("user_login") != "" {
		err := user.User.UpdateUserByID(id, "user_login", r.URL.Query().Get("user_login"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if r.URL.Query().Get("user_pass") != "" {
		err := user.User.UpdateUserByID(id, "user_pass", r.URL.Query().Get("user_pass"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if r.URL.Query().Get("user_finapi_pass") != "" {
		err := user.User.UpdateUserByID(id, "user_finapi_pass", r.URL.Query().Get("user_finapi_pass"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if r.URL.Query().Get("user_nicename") != "" {
		err := user.User.UpdateUserByID(id, "user_nicename", r.URL.Query().Get("user_nicename"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if r.URL.Query().Get("user_email") != "" {
		err := user.User.UpdateUserByID(id, "user_email", r.URL.Query().Get("user_email"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if r.URL.Query().Get("activation_code") != "" {
		err := user.User.UpdateUserByID(id, "activation_code", r.URL.Query().Get("activation_code"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if r.URL.Query().Get("user_registered") != "" {
		err := user.User.UpdateUserByID(id, "user_registered", r.URL.Query().Get("user_registered"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if r.URL.Query().Get("is_verified") != "" {
		err := user.User.UpdateUserByID(id, "is_verified", r.URL.Query().Get("is_verified"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if r.URL.Query().Get("last_name") != "" {
		err := user.User.UpdateUserByID(id, "last_name", r.URL.Query().Get("last_name"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if r.URL.Query().Get("first_name") != "" {
		err := user.User.UpdateUserByID(id, "first_name", r.URL.Query().Get("first_name"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	if r.URL.Query().Get("user_credentials") != "" {
		err := user.User.UpdateUserByID(id, "user_credentials", r.URL.Query().Get("user_credentials"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	return
}

func (user *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := user.User.GetUsers()
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
