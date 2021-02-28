package commands

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/user-service/api/adapters"
	"github.com/thomas-bousquet/user-service/errors"
	. "github.com/thomas-bousquet/user-service/models"
	. "github.com/thomas-bousquet/user-service/repositories"
	. "github.com/thomas-bousquet/user-service/utils/jwt"
	. "github.com/thomas-bousquet/user-service/utils/validator"
	"golang.org/x/crypto/bcrypt"
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

func (c LoginCommand) Execute(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) *errors.Error {
	email, password, ok := r.BasicAuth()
	defaultErrorMessage := "An error occurred when logging user"

	if !ok {
		return errors.NewAuthorizationError(defaultErrorMessage)
	}

	credentials := NewCredentials(email, password)
	validationErrors := c.validator.ValidateStruct(credentials)

	if len(validationErrors) > 0 {
		return errors.NewValidationError(
			"An error occurred when validating credentials",
			validationErrors,
		)
	}

	user, err := c.userRepository.FindUserByEmail(credentials.Email)

	if err != nil {
		logger.Error(err)
		return errors.NewUnexpectedError()
	}

	if user == nil {
		return errors.NewAuthorizationError(defaultErrorMessage)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		logger.Errorf("error when comparing passwords: %v", err)
		return errors.NewAuthorizationError(defaultErrorMessage)
	}

	token, err := c.jwt.CreateToken(*user)

	if err != nil {
		logger.Errorf("error creating JWT token: %v", err)
		return errors.NewUnexpectedError()
	}

	response, err := json.Marshal(NewLoginAdapter(*token))

	if err != nil {
		logger.Errorf("error marshalling response: %v", err)
		return errors.NewUnexpectedError()
	}

	_, err = w.Write(response)

	if err != nil {
		logger.Errorf("error writing response %v", err)
		return errors.NewUnexpectedError()
	}

	return nil
}
