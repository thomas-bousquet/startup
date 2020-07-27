package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/thomas-bousquet/startup/repository"
	"net/http"
)

type ReadUserHandler struct {
	userRepository UserRepository
}

func NewReadUserHandler(userRepository UserRepository) CreateUserHandler {
	return CreateUserHandler {
		userRepository,
	}
}

func (h ReadUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	println(email)
	println(email)
	println(email)
	println(email)
	println(email)

	user := h.userRepository.FindUserByEmail(email)

	json.NewEncoder(w).Encode(user)
}