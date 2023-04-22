/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package services

import (
	"log"
	"test/database"
	"test/models"
)

type GroupService struct {
	repository database.GroupRepository
}

func NewGroupService(repo database.GroupRepository) *GroupService {
	return &GroupService{repository: repo}
}

func (service *GroupService) CreateGroup(name string,
	status string,
	owner string,
) ([]models.Group, error) {
	var group []models.Group

	group, err := service.repository.CreateGroup(name,
		status,
		owner)

	if err != nil {
		log.Println(err)
		return nil, err
	}

	var newGroup []models.Group
	for _, tmpGroup := range group {
		newGroup = append(newGroup, models.Group{
			ID:          tmpGroup.ID,
			Name:        tmpGroup.Name,
			GroupStatus: tmpGroup.GroupStatus,
			GroupOwner:  tmpGroup.GroupOwner,
		})
	}

	return newGroup, nil
}

func (service *GroupService) GetGroups() ([]models.Group, error) {
	var group []models.Group

	group, err := service.repository.GetGroups()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var newGroup []models.Group
	for _, tmpGroup := range group {
		newGroup = append(newGroup, models.Group{
			ID:          tmpGroup.ID,
			Name:        tmpGroup.Name,
			GroupStatus: tmpGroup.GroupStatus,
			GroupOwner:  tmpGroup.GroupOwner,
		})
	}

	return newGroup, nil
}
