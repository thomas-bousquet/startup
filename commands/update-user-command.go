package commands

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/startup/errors"
	. "github.com/thomas-bousquet/startup/models"
	. "github.com/thomas-bousquet/startup/repositories"
	"github.com/thomas-bousquet/startup/utils/validator"
	"net/http"
)

type UpdateUserCommand struct {
	userRepository UserRepository
	validator      validator.Validator
}

func NewUpdateUserCommand(userRepository UserRepository, validator validator.Validator) UpdateUserCommand {
	return UpdateUserCommand{
		userRepository,
		validator,
	}
}

func (c UpdateUserCommand) Execute(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) error {
	logger.Info("Update user")
	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		return err
	}

	user.Id = id

	errors := c.validator.ValidateStructExcept(user, "Password")

	if len(errors) > 0 {
		return NewValidationError("An error occurred when validating user fields", errors)
	}

	err = c.userRepository.UpdateUser(id, user)

	if err != nil {
		return err
	}

	return nil
}