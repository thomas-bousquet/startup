package api

import (
	"github.com/gorilla/mux"
	. "github.com/thomas-bousquet/startup/command"
	. "github.com/thomas-bousquet/startup/repository"
	"github.com/thomas-bousquet/startup/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func RegisterRoutes(router *mux.Router, mongoDB *mongo.Database) {
	userRepository := NewUserRepository(mongoDB.Collection("users"))
	customValidator := utils.NewValidator(validator.New())
	createUserHandler := NewHandler(NewCreateUserCommand(userRepository, customValidator))
	updateUserHandler := NewHandler(NewUpdateUserCommand(userRepository, customValidator))
	readUserHandler := NewHandler(NewReadUserCommand(userRepository))
	readUserByEmailHandler := NewHandler(NewReadUserByEmailCommand(userRepository))

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})
	router.Handle("/users", createUserHandler).Methods("POST")
	router.Handle("/users/{id}", updateUserHandler).Methods("PUT")
	router.Handle("/users/{id}", readUserHandler).Methods("GET")
	router.Handle("/users/email/{email}", readUserByEmailHandler).Methods("GET")
}