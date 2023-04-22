package database

import "test/models"

type Group struct {
	ID          string `gorm:"ID;PRIMARY_KEY;BIGINT(20);AUTO_INCREMENT;NOT NULL"`
	Name        string `gorm:"name;VARCHAR(250);NULL"`
	GroupStatus string `gorm:"group_status;VARCHAR(1);NOT NULL"`
	GroupOwner  string `gorm:"group_owner;INT;NULL"`
}

type GroupRepository interface {
	GetGroups() ([]models.Group, error)
	CreateGroup(
		name string,
		status string,
		owner string,
	) ([]models.Group, error)
}

func (*Group) TableName() string {
	return "paygoal_app.user_group"
}

func (db *DB) CreateGroup(name string, status string, owner string) ([]models.Group, error) {
	var group []Group

	err := db.db.Table("paygoal_app.user_group").Create(&models.Group{
		ID:          "",
		Name:        name,
		GroupStatus: status,
		GroupOwner:  owner,
	}).Error

	if err != nil {
		return nil, err
	}

	var newGroup []models.Group
	for _, tmpGroup := range group {
		newGroup = append(newGroup, models.Group{
			ID:          "",
			Name:        tmpGroup.Name,
			GroupStatus: tmpGroup.GroupStatus,
			GroupOwner:  tmpGroup.GroupOwner,
		})
	}

	return newGroup, nil
}

func (db *DB) GetGroups() ([]models.Group, error) {
	var group []Group

	err := db.db.
		Select(
			"paygoal_app.user_group.ID as ID",
			"paygoal_app.user_group.name as name",
			"paygoal_app.user_group.group_status as group_status",
			"paygoal_app.user_group.group_owner as group_owner").
		Find(&group).Error

	if err != nil {
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
