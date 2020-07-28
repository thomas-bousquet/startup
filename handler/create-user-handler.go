package handler

import (
	"encoding/json"
	. "github.com/thomas-bousquet/startup/model"
	. "github.com/thomas-bousquet/startup/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	var user = User{}
	json.NewDecoder(r.Body).Decode(&user)
	userId := h.userRepository.CreateUser(user)

	response, _ := json.Marshal(map[string]primitive.ObjectID{"id": userId})
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}