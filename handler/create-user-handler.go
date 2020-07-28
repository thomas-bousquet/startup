package handler

import (
	"encoding/json"
	. "github.com/thomas-bousquet/startup/model"
	. "github.com/thomas-bousquet/startup/repository"
	"net/http"
)

type CreateUserHandler struct {
	userRepository UserRepository
}

func NewCreateUserHandler(userRepository UserRepository) CreateUserHandler {
	return CreateUserHandler {
		userRepository,
	}
}

func (h CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var userPayload = User{}
	json.NewDecoder(r.Body).Decode(&userPayload)

	user := NewUser(userPayload.Firstname, userPayload.Lastname, userPayload.Email, userPayload.Password)
	h.userRepository.CreateUser(user)
}