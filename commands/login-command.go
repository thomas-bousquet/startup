package commands

import (
	"encoding/json"
	. "startup/api/adapters"
	. "startup/errors"
	. "startup/models"
	. "startup/repositories"
	. "startup/utils/jwt"
	. "startup/utils/validator"
	"net/http"
)

type LoginCommand struct {
	userRepository UserRepository
	validator      Validator
	jwt            JWT
}

func NewLoginCommand(userRepository UserRepository, validator Validator, jwt JWT) LoginCommand {
	return LoginCommand{
		userRepository: userRepository,
		validator:      validator,
		jwt:            jwt,
	}
}

func (c LoginCommand) Execute(w http.ResponseWriter, r *http.Request) error {
	email, password, ok := r.BasicAuth()
	defaultErrorMessage := "An error occurred when logging user in"

	if !ok {
		return NewAuthorizationError(defaultErrorMessage)
	}

	credentials := NewCredentials(email, password)
	errors := c.validator.ValidateStruct(credentials)

	if len(errors) > 0 {
		return NewValidationError(
			"An error occurred when validation credentials",
			errors,
		)
	}

	user, err := c.userRepository.AuthenticateUser(credentials.Email, credentials.Password)

	if err != nil {
		return NewUnexpectedError()
	}

	if user == nil {
		return NewAuthorizationError(defaultErrorMessage)
	}

	token, err := c.jwt.CreateToken(*user)

	if err != nil {
		return err
	}

	response, err := json.Marshal(NewLoginAdapter(*token))

	_, err = w.Write(response)

	return err
}
