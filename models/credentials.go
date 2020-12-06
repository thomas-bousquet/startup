package models

type Credentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func NewCredentials(email string, password string) Credentials {
	return Credentials{Email: email, Password: password}
}
