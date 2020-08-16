package adapter

import (
	"github.com/thomas-bousquet/startup/model"
)

type UserAdapter struct {
	Id string `json:"id"`
	Firstname string `json:"first_name"`
	Lastname string `json:"last_name"`
	Email string `json:"email"`
}

func NewUserAdapter(user *model.User) UserAdapter {
	return UserAdapter{
		Id: user.Id,
		Firstname: user.Firstname,
		Lastname:  user.Lastname,
		Email:     user.Email,
	}
}