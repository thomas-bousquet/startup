package commands

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/startup/errors"
	. "github.com/thomas-bousquet/startup/models"
	. "github.com/thomas-bousquet/startup/repositories"
	"github.com/thomas-bousquet/startup/utils/validator"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type CreateUserCommand struct {
	userRepository UserRepository
	validator      validator.Validator
}

func NewCreateUserCommand(userRepository UserRepository, validator validator.Validator) CreateUserCommand {
	return CreateUserCommand{
		userRepository: userRepository,
		validator:      validator,
	}
}

func (c CreateUserCommand) Execute(w http.ResponseWriter, r *http.Request) error {
	log.Info("Creating user")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		return err
	}

	errors := c.validator.ValidateStruct(user)

	if len(errors) > 0 {
		return NewValidationError("An error occurred when validating user fields", errors)
	}

	userId, err := c.userRepository.CreateUser(user)

	if err != nil {
		return err
	}

	response, err := json.Marshal(map[string]primitive.ObjectID{"id": userId})

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	 _, err = w.Write(response)

	return err
}
