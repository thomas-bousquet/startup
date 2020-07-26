package api

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	. "github.com/thomas-bousquet/startup/handler"
	. "github.com/thomas-bousquet/startup/repository"
)

func RegisterRoutes(router *mux.Router, mongoClient *mongo.Client) {
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})

	userRepository := NewUserRepository(mongoClient)
	createUserHandler := NewCreateUserHandler(userRepository)

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})
	router.HandleFunc("/users", createUserHandler.Handle).Methods("POST")
}
