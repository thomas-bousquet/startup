package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/thomas-bousquet/startup/repository"
	"net/http"
)

type ReadUserByEmailHandler struct {
	userRepository UserRepository
}

func NewReadUserByEmailHandler(userRepository UserRepository) ReadUserByEmailHandler {
	return ReadUserByEmailHandler{
		userRepository,
	}
}

func (h ReadUserByEmailHandler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	email := vars["email"]

	user := h.userRepository.FindUserByEmail(email)

	json.NewEncoder(w).Encode(user)
}