package handler

import (
	"github.com/gorilla/mux"
	"github.com/thomas-bousquet/startup/model"
	. "github.com/thomas-bousquet/startup/repository"
	"net/http"
	"time"
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
	vars := mux.Vars(r)
	println(vars)
	user := model.User{
		Firstname: "John",
		Lastname:  "Does",
		Email:     "john.doe@gmail.com",
		Password:  "12345",
		CreatedAt: time.Now(),
	}
	h.userRepository.CreateUser(user)
}