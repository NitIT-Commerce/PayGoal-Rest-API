/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package services

import (
	_ "github.com/go-sql-driver/mysql"
	"test/models"
	"test/utils"
)

type UserService struct {
	container *utils.Helper
}

func NewUserService(container *utils.Helper) *UserService {
	return &UserService{container: container}
}

func (user *UserService) GetAllUsers() ([]models.Users, error) {
	var users []models.Users

	/*, user_login, user_pass, user_finapi_pass, user_nicename, user_email, activation_code, user_registered, is_verified, last_name, first_name, user_credentials" */
	err := user.container.DB.Select("paygoal_app.users.ID").First(users).Error
	if err != nil {
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
}
