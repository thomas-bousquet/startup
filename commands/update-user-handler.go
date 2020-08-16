package commands

import (
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/thomas-bousquet/startup/models"
	. "github.com/thomas-bousquet/startup/repositories"
	"github.com/thomas-bousquet/startup/utils/validator"
	. "github.com/thomas-bousquet/startup/errors"
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

func (h UpdateUserCommand) Execute(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		return err
	}

	user.Id = id

	errors := h.validator.ValidateStructExcept(user, "Password")

	if len(errors) > 0 {
		return NewValidationError(errors)
	}

	err = h.userRepository.UpdateUser(id, user)

	if err != nil {
		return err
	}

	return nil
}