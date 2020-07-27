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

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("/users", createUserHandler.Handle).Methods("POST")
}
