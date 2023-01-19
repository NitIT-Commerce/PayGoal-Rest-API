package models

type Users struct {
	ID string `json:"id"`
	/*UserLogin       string `json:"user_login"`
	UserPass        string `json:"user_pass"`
	UserFinApiPass  string `json:"user_finapi_pass"`
	UserNicename    string `json:"user_nicename"`
	UserEmail       string `json:"user_email"`
	ActivationCode  int    `json:"activation_code"`
	UserRegistered  string `json:"user_registered"`
	IsVerified      bool   `json:"is_verified"`
	LastName        string `json:"last_name"`
	FirstName       string `json:"first_name"`
	UserCredentials int    `json:"user_credentials"`*/
}

type UserService interface {
	GetAllUsers() ([]Users, error)
}

func (*Users) TableName() string {
	return "users"
}
