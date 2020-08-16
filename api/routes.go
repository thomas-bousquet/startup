package api

import (
	"github.com/gorilla/mux"
	. "github.com/thomas-bousquet/startup/commands"
	. "github.com/thomas-bousquet/startup/repositories"
	. "github.com/thomas-bousquet/startup/utils/validator"
	JWT "github.com/thomas-bousquet/startup/utils/jwt"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func RegisterRoutes(router *mux.Router, mongoDB *mongo.Database) {
	userRepository := NewUserRepository(mongoDB.Collection("users"))
	customValidator := NewValidator(validator.New())
	jwt := JWT.New([]byte("my_secret_key"))

	createUserHandler := NewHandler(NewCreateUserCommand(userRepository, customValidator))
	updateUserHandler := NewHandler(NewUpdateUserCommand(userRepository, customValidator))
	readUserHandler := NewHandler(NewReadUserCommand(userRepository))
	readUserByEmailHandler := NewHandler(NewReadUserByEmailCommand(userRepository))
	loginHandler := NewHandler(NewLoginCommand(userRepository, customValidator, jwt))

	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})
	router.Handle("/users", createUserHandler).Methods("POST")
	router.Handle("/users/{id}", updateUserHandler).Methods("PUT")
	router.Handle("/users/{id}", readUserHandler).Methods("GET")
	router.Handle("/users/email/{email}", readUserByEmailHandler).Methods("GET")
	router.Handle("/login", loginHandler).Methods("POST")
}
