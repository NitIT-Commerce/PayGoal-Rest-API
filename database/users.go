/*
 * Copyright (c) by Lukas Nickel, PayGoal UG 2023.
 */

package database

import (
	"test/models"
)

type Users struct {
	ID              string `gorm:"ID;BIGINT(20)"`
	UserLogin       string `gorm:"user_login;VARCHAR(60)"`
	UserPass        string `gorm:"user_pass;VARCHAR(255)"`
	UserFinApiPass  string `gorm:"user_finapi_pass;VARCHAR(500)"`
	UserNicename    string `gorm:"user_nicename;VARCHAR(50)"`
	UserEmail       string `gorm:"user_email;VARCHAR(100)"`
	ActivationCode  string `gorm:"activation_code;VARCHAR(255)"`
	UserRegistered  string `gorm:"user_registered;DATETIME"`
	IsVerified      int    `gorm:"is_verified;TINYINT(1)"`
	LastName        string `gorm:"last_name;VARCHAR(250)"`
	FirstName       string `gorm:"first_name;VARCHAR(250)"`
	UserCredentials string `gorm:"user_credentials;VARCHAR(250)"`
}

type UserRepository interface {
	GetUsers() ([]models.Users, error)
}

func (*Users) TableName() string {
	return "paygoal_app.users"
}

func (db *DB) GetUsers() ([]models.Users, error) {
	var users []Users
	err := db.db.
		Select("paygoal_app.users.ID as ID, " +
			"paygoal_app.users.user_login as user_login, " +
			"paygoal_app.users.user_pass, " +
			"paygoal_app.users.user_finapi_pass, " +
			"paygoal_app.users.user_nicename, " +
			"paygoal_app.users.user_email, " +
			"paygoal_app.users.activation_code," +
			"paygoal_app.users.user_registered," +
			"paygoal_app.users.is_verified," +
			"paygoal_app.users.last_name," +
			"paygoal_app.users.first_name," +
			"paygoal_app.users.user_credentials ").
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	var newUsers []models.Users
	for _, user := range users {
		newUsers = append(newUsers, models.Users{
			ID:              user.ID,
			UserLogin:       user.UserLogin,
			UserPass:        user.UserPass,
			UserFinApiPass:  user.UserFinApiPass,
			UserNicename:    user.UserNicename,
			UserEmail:       user.UserEmail,
			ActivationCode:  user.ActivationCode,
			UserRegistered:  user.UserRegistered,
			IsVerified:      user.IsVerified,
			LastName:        user.LastName,
			FirstName:       user.FirstName,
			UserCredentials: user.UserCredentials,
		})
	}

	return newUsers, err
}
