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

func NewReadUserHandler(userRepository UserRepository) ReadUserHandler {
	return ReadUserHandler {
		userRepository,
	}
}

func (h ReadUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user := h.userRepository.FindUser(id)

	json.NewEncoder(w).Encode(user)
}