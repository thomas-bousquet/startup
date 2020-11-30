package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/startup/commands"
	"github.com/thomas-bousquet/startup/middlewares"
	. "github.com/thomas-bousquet/startup/repositories"
	JWT "github.com/thomas-bousquet/startup/utils/jwt"
	. "github.com/thomas-bousquet/startup/utils/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func RegisterRoutes(router *mux.Router, mongoClient *mongo.Client, logger *logrus.Logger) {
	// Repositories
	userRepository := NewUserRepository(mongoClient, logger)
	customValidator := NewValidator(validator.New())
	jwt := JWT.New([]byte("my_secret_key"))

	// Middlewares
	authenticationMiddleware := middlewares.NewAuthenticationMiddleware(jwt, userRepository, logger)

	// Commands
	createUserHandler := NewHandler(NewCreateUserCommand(userRepository, customValidator), logger)
	updateUserHandler := NewHandler(NewUpdateUserCommand(userRepository, customValidator), logger)
	readUserHandler := NewHandler(NewReadUserCommand(userRepository), logger)
	readUsersHandler := NewHandler(NewReadUsersCommand(userRepository), logger)
	readUserByEmailHandler := NewHandler(NewReadUserByEmailCommand(userRepository), logger)
	loginHandler := NewHandler(NewLoginCommand(userRepository, customValidator, jwt), logger)

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
