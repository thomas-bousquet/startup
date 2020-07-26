package handler

import (
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
}