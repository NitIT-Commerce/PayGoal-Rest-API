package models

type CD struct {
	ID string `json:"ID"`
}

type testRepository interface {
	GetUsers() ([]CD, error)
}
