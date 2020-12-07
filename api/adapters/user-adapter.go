package adapters

import (
	"github.com/thomas-bousquet/user-service/models"
)

type UserAdapter struct {
	Id        string `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
	Email     string `json:"email"`
}

func NewUserAdapter(user *models.User) UserAdapter {
	return UserAdapter{
		Id:        user.Id,
		Firstname: user.FirstName,
		Lastname:  user.LastName,
		Email:     user.Email,
	}
}
