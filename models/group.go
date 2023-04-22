package models

type Group struct {
	ID          string `json:"ID"`
	Name        string `json:"name"`
	GroupStatus string `json:"group_status"`
	GroupOwner  string `json:"group_owner"`
}

type GroupService interface {
	CreateGroup(name string, status string, owner string) ([]Group, error)
	GetGroups() ([]Group, error)
}
