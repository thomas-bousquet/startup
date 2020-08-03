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
	router.HandleFunc("/users", createUserHandler.Handle).Methods("POST")
	router.HandleFunc("/users/{id}", updateUserHandler.Handle).Methods("PUT")
	router.HandleFunc("/users/{id}", readUserHandler.Handle).Methods("GET")
	router.HandleFunc("/users/email/{email}", readUserByEmailHandler.Handle).Methods("GET")
}
