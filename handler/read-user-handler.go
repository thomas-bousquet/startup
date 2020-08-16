package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/thomas-bousquet/startup/api/adapter"
	. "github.com/thomas-bousquet/startup/repository"
	"net/http"
)

type ReadUserHandler struct {
	userRepository UserRepository
}

func NewReadUserHandler(userRepository UserRepository) ReadUserHandler {
	return ReadUserHandler{
		userRepository,
	}
}

func (h ReadUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user := h.userRepository.FindUser(id)

	response, _ := json.Marshal(adapter.NewUserAdapter(user))
	w.Write(response)
}
