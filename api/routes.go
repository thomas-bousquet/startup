package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/user-service/commands"
	"github.com/thomas-bousquet/user-service/middlewares"
	. "github.com/thomas-bousquet/user-service/repositories"
	. "github.com/thomas-bousquet/user-service/utils/error-handler"
	JWT "github.com/thomas-bousquet/user-service/utils/jwt"
	. "github.com/thomas-bousquet/user-service/utils/validator"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

func RegisterRoutes(router *mux.Router, mongoClient *mongo.Client, logger *logrus.Logger) {
	// Repositories
	userRepository := NewUserRepository(mongoClient, logger)
	customValidator := NewValidator(validator.New())
	jwt := JWT.New([]byte("my_secret_key"))
	errorHandler := NewErrorHandler()

	// Middlewares
	authenticationMiddleware := middlewares.NewAuthenticationMiddleware(jwt, userRepository, logger, errorHandler)

	// Commands
	createUserHandler := NewHandler(NewCreateUserCommand(userRepository, customValidator), logger, errorHandler)
	updateUserHandler := NewHandler(NewUpdateUserCommand(userRepository, customValidator), logger, errorHandler)
	readUserHandler := NewHandler(NewReadUserCommand(userRepository), logger, errorHandler)
	readUsersHandler := NewHandler(NewReadUsersCommand(userRepository), logger, errorHandler)
	//readUserByEmailHandler := NewHandler(NewReadUserByEmailCommand(userRepository), logger, errorHandler)
	loginHandler := NewHandler(NewLoginCommand(userRepository, customValidator, jwt), logger, errorHandler)

	// Router
	router.HandleFunc("/admin/health", func(w http.ResponseWriter, r *http.Request) {})
	router.Handle("/login", loginHandler).Methods("POST")
	router.Handle("/users", createUserHandler).Methods("POST")

	// UserAuthRouter
	userRouter := router.PathPrefix("/users").Subrouter()
	userRouter.Use(authenticationMiddleware.ExecuteWithRole("user"))
	userRouter.Handle("/{id}", updateUserHandler).Methods("PUT")
	userRouter.Handle("/{id}", readUserHandler).Methods("GET")

	// Admin Router
	adminRouter := router.PathPrefix("/admin").Subrouter()
	adminRouter.Use(authenticationMiddleware.ExecuteWithRole("admin"))
	adminRouter.Handle("/users", readUsersHandler).Methods("GET")

	//userRouter.Handle("/email/{email}", readUserByEmailHandler).Methods("GET")
}
