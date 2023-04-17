package models

type Users struct {
	ID              string `json:"ID"`
	UserLogin       string `json:"user_login"`
	UserPass        string `json:"user_pass"`
	UserFinApiPass  string `json:"user_fin_api_pass"`
	UserNicename    string `json:"user_nicename"`
	UserEmail       string `json:"user_email"`
	ActivationCode  string `json:"activation_code"`
	UserRegistered  string `json:"user_registered"`
	IsVerified      int    `json:"is_verified"`
	LastName        string `json:"last_name"`
	FirstName       string `json:"first_name"`
	UserCredentials string `json:"user_credentials"`
}

type UserService interface {
	GetUsers() ([]Users, error)
	GetUserByID(userid string) ([]Users, error)
	UpdateUserByID(userid string, field string, value string) error
	DeleteUserByID(userid string) error
	CreateUser(userEmail string,
		userPass string,
		userNickname string,
		userName string,
		userLastName string) ([]Users, error)
}
