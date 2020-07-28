package handler

import (
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/thomas-bousquet/startup/model"
	. "github.com/thomas-bousquet/startup/repository"
	"net/http"
)

type UpdateUserHandler struct {
	userRepository UserRepository
}

func NewUpdateUserHandler(userRepository UserRepository) UpdateUserHandler {
	return UpdateUserHandler {
		userRepository,
	}
}

func (h UpdateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	updateUserPayload := User{}
	_ = json.NewDecoder(r.Body).Decode(&updateUserPayload)
	updateUserPayload.Id = id

	h.userRepository.UpdateUser(id, updateUserPayload)

	w.WriteHeader(http.StatusNoContent)
}