package api

import (
	"github.com/gorilla/mux"
	. "github.com/thomas-bousquet/startup/handler"
	. "github.com/thomas-bousquet/startup/repository"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

func RegisterRoutes(router *mux.Router, mongoDB *mongo.Database) {
	userRepository := NewUserRepository(mongoDB.Collection("users"))
	createUserHandler := NewCreateUserHandler(userRepository)
	readUserHandler := NewReadUserHandler(userRepository)
	readUserByEmailHandler := NewReadUserHandler(userRepository)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("/users", createUserHandler.Handle).Methods("POST")
	router.HandleFunc("/users/id/{id}", readUserHandler.Handle).Methods("GET")
	router.HandleFunc("/users/email/{email}", readUserByEmailHandler.Handle).Methods("GET")
}
