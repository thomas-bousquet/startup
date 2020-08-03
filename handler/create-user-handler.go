package handler

import (
	"encoding/json"
	"fmt"
	. "github.com/thomas-bousquet/startup/model"
	. "github.com/thomas-bousquet/startup/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type CreateUserHandler struct {
	userRepository UserRepository
	validator      *validator.Validate
}

func NewCreateUserHandler(userRepository UserRepository, validator *validator.Validate) CreateUserHandler {
	return CreateUserHandler{
		userRepository,
		validator,
	}
}

func (h CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var user = User{}
	json.NewDecoder(r.Body).Decode(&user)

	err := h.validator.Struct(user)

	for _, e := range err.(validator.ValidationErrors) {
		fmt.Println(e)
	}

	userId := h.userRepository.CreateUser(user)

	response, _ := json.Marshal(map[string]primitive.ObjectID{"id": userId})
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
