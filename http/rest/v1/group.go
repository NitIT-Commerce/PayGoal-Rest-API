/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package v1

import (
	"encoding/json"
	"net/http"
	"test/services"
)

type GroupController struct {
	Group *services.GroupService
}

func NewGroupController(group *services.GroupService) *GroupController {
	return &GroupController{
		Group: group,
	}
}

func (group *GroupController) GetGroups(w http.ResponseWriter, r *http.Request) {
	groups, err := group.Group.GetGroups()
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	marshal, err := json.Marshal(groups)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
	w.WriteHeader(http.StatusOK)
	return
}

func (group *GroupController) CreateGroup(w http.ResponseWriter, r *http.Request) {
	groups, err := group.Group.CreateGroup(
		r.URL.Query().Get("name"),
		r.URL.Query().Get("group_status"),
		r.URL.Query().Get("group_owner"),
	)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	marshal, err := json.Marshal(groups)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(marshal)
	w.WriteHeader(http.StatusOK)
	return
}
