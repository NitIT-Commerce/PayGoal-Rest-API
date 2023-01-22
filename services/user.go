/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package services

import (
	_ "github.com/go-sql-driver/mysql"
	"log"
	"test/models"
)

type UserService struct {
	repository models.UserRepository
}

func NewUserService(repo models.UserRepository) *UserService {
	return &UserService{repository: repo}
}

func (service *UserService) GetAllUsers() ([]models.CD, error) {
	var users []models.CD

	/*, user_login, user_pass, user_finapi_pass, user_nicename, user_email, activation_code, user_registered, is_verified, last_name, first_name, user_credentials" */
	users, err := service.repository.GetUsers()
	if err != nil {
		log.Println("Somethig went wrong")
		return nil, err
	}

	var newUsers []models.Users
	for _, user := range users {
		newUsers = append(newUsers, models.Users{
			ID: user.ID,
			/*UserLogin:       user.UserLogin,
			UserPass:        user.UserPass,
			UserFinApiPass:  user.UserFinApiPass,
			UserNicename:    user.UserNicename,
			UserEmail:       user.UserEmail,
			ActivationCode:  user.ActivationCode,
			UserRegistered:  user.UserRegistered,
			IsVerified:      user.IsVerified,
			LastName:        user.LastName,
			FirstName:       user.FirstName,
			UserCredentials: user.UserCredentials,*/
		})
	}

	return newUsers, nil

	/*var newUsers []models.Users
	newUsers = append(newUsers, models.Users{
		ID: "1",
	})

	return newUsers, nil*/
}
