package commands

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/user-service/app_errors"
	. "github.com/thomas-bousquet/user-service/models"
	. "github.com/thomas-bousquet/user-service/repositories"
	"github.com/thomas-bousquet/user-service/utils/validator"
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
	vars := mux.Vars(r)
	id := vars["id"]

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		logger.Errorf("error unmarshalling request: %v", err)
		return app_errors.NewUnexpectedError(nil, nil)
	}

	user.Id = id

	validationErrors := c.validator.ValidateStructExcept(user, "Password")

	if len(validationErrors) > 0 {
		return app_errors.NewValidationError("An error occurred when validating user fields", validationErrors)
	}

	err = c.userRepository.UpdateUser(id, user)

	if err != nil {
		logger.Errorf("error updating user: %v", err)
		return app_errors.NewUnexpectedError(nil, nil)
	}

	return nil
}
