package database

import "test/models"

type Users struct {
	ID string `json:"ID"`
}

func (*Users) TableName() string {
	return "paygoal_app.users"
}

func (db *DB) GetUsers() ([]models.Users, error) {
	var users []Users
	err := db.db.
		Select("paygoal_app.users.ID").
		Find(&users).Error
	if err != nil {
		return nil, err
	}

	var newUsers []models.Users
	for _, user := range users {
		newUsers = append(newUsers, models.Users{
			ID: user.ID,
		})
	}
	return newUsers, nil
}
