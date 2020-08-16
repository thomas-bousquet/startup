package api

import (
	"github.com/gorilla/mux"
	. "github.com/thomas-bousquet/startup/handler"
	. "github.com/thomas-bousquet/startup/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func RegisterRoutes(router *mux.Router, mongoDB *mongo.Database) {
	userRepository := NewUserRepository(mongoDB.Collection("users"))
	validate := validator.New()
	createUserHandler := NewCreateUserHandler(userRepository, validate)
	updateUserHandler := NewUpdateUserHandler(userRepository, validate)
	readUserHandler := NewReadUserHandler(userRepository)
	readUserByEmailHandler := NewReadUserByEmailHandler(userRepository)


	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})
	router.Handle("/users", createUserHandler).Methods("POST")
	router.Handle("/users/{id}", updateUserHandler).Methods("PUT")
	router.Handle("/users/{id}", readUserHandler).Methods("GET")
	router.Handle("/users/email/{email}", readUserByEmailHandler).Methods("GET")
}
