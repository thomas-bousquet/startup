package commands

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	. "startup/errors"
	. "startup/models"
	. "startup/repositories"
	"startup/utils/validator"
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

func (c UpdateUserCommand) Execute(w http.ResponseWriter, r *http.Request) error {
	log.Info("Update user")
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