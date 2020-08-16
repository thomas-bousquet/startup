package commands

import (
	"encoding/json"
	. "github.com/thomas-bousquet/startup/api/adapters"
	. "github.com/thomas-bousquet/startup/errors"
	. "github.com/thomas-bousquet/startup/models"
	. "github.com/thomas-bousquet/startup/repositories"
	. "github.com/thomas-bousquet/startup/utils/jwt"
	. "github.com/thomas-bousquet/startup/utils/validator"
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

	if !ok {
		return NewAuthenticationError()
	}

	credentials := NewCredentials(email, password)
	errors := c.validator.ValidateStruct(credentials)

	if len(errors) > 0 {
		return NewValidationError(errors)
	}

	user, err := c.userRepository.AuthenticateUser(credentials.Email, credentials.Password)

	if err != nil {
		return NewUnexpectedError()
	}

	if user == nil {
		return NewAuthenticationError()
	}

	token, err := c.jwt.CreateToken(*user)

	if err != nil {
		return err
	}

	response, err := json.Marshal(NewLoginAdapter(*token))

	_, err = w.Write(response)

	return err
}
