package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	. "github.com/thomas-bousquet/startup/model"
	. "github.com/thomas-bousquet/startup/repository"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type UpdateUserHandler struct {
	userRepository UserRepository
	validator      *validator.Validate
}

func NewUpdateUserHandler(userRepository UserRepository, validator *validator.Validate) UpdateUserHandler {
	return UpdateUserHandler {
		userRepository,
		validator,
	}
}

func (h UpdateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	user := User{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	user.Id = id

	err := h.validator.StructExcept(user, "password")
	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e)
	}

	h.userRepository.UpdateUser(id, user)

	w.WriteHeader(http.StatusNoContent)
}