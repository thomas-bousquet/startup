package model

import (
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id string `json:"id" bson:"id"`
	Firstname string `json:"firstname" bson:"firstname"`
	Lastname string `json:"lastname" bson:"lastname"`
	Email string `json:"email" bson:"email"`
	Password string `json:"-" bson:"password"`
	CreatedAt time.Time `json:"-" bson:"created_at"`
	UpdatedAt time.Time `json:"-" bson:"updated_at"`
	DeletedAt time.Time `json:"-" bson:"deleted_at"`
}

func NewUser(firstname string, lastname string, email string, password string) User {
	return User{
		Id: uuid.New().String(),
		Firstname: firstname,
		Lastname:  lastname,
		Email:     email,
		Password: password,
		CreatedAt: time.Now(),
	}
}
