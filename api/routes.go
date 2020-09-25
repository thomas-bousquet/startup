package api

import (
	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
	. "startup/commands"
	"startup/middlewares"
	. "startup/repositories"
	JWT "startup/utils/jwt"
	. "startup/utils/validator"
)

func RegisterRoutes(router *mux.Router, mongoDB *mongo.Database) {
	// Repositories
	userRepository := NewUserRepository(mongoDB.Collection("users"))
	customValidator := NewValidator(validator.New())
	jwt := JWT.New([]byte("my_secret_key"))

	// Middlewares
	authenticationMiddleware := middlewares.NewAuthenticationMiddleware(jwt, userRepository)

	// Commands
	createUserHandler := NewHandler(NewCreateUserCommand(userRepository, customValidator))
	updateUserHandler := NewHandler(NewUpdateUserCommand(userRepository, customValidator))
	readUserHandler := NewHandler(NewReadUserCommand(userRepository))
	readUsersHandler := NewHandler(NewReadUsersCommand(userRepository))
	readUserByEmailHandler := NewHandler(NewReadUserByEmailCommand(userRepository))
	loginHandler := NewHandler(NewLoginCommand(userRepository, customValidator, jwt))

	// Router
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {})
	router.Handle("/login", loginHandler).Methods("POST")
	router.Handle("/users", createUserHandler).Methods("POST")

	// Admin Router
	adminRouter := router.PathPrefix("").Subrouter()
	adminRouter.Use(authenticationMiddleware.ExecuteWithRole("admin"))
	adminRouter.Handle("/users", readUsersHandler).Methods("GET")

	// User Router
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.Use(authenticationMiddleware.ExecuteWithRole("user"))

	userRouter.Handle("/{id}", updateUserHandler).Methods("PUT")
	userRouter.Handle("/{id}", readUserHandler).Methods("GET")
	userRouter.Handle("/email/{email}", readUserByEmailHandler).Methods("GET")
}
