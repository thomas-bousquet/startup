package commands

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/user-service/api/adapters"
	"github.com/thomas-bousquet/user-service/app_errors"
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

func (c LoginCommand) Execute(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) error {
	email, password, ok := r.BasicAuth()

	if !ok {
		return app_errors.NewAuthorizationError(nil)
	}

	credentials := NewCredentials(email, password)
	validationErrors := c.validator.ValidateStruct(credentials)

	if len(validationErrors) > 0 {
		return app_errors.NewValidationError(
			"An error occurred when validating credentials",
			validationErrors,
		)
	}

	user, err := c.userRepository.FindUserByEmail(credentials.Email)

	if err != nil {
		logger.Errorf("%v", err)
		return app_errors.NewUnexpectedError(nil, nil)
	}

	if user == nil {
		return app_errors.NewAuthorizationError(nil)
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		logger.Errorf("error when comparing passwords: %v", err)
		return app_errors.NewAuthorizationError(nil)
	}

	token, err := c.jwt.CreateToken(*user)

	if err != nil {
		logger.Errorf("error creating JWT token: %v", err)
		return app_errors.NewUnexpectedError(nil, nil)
	}

	response, err := json.Marshal(NewLoginAdapter(*token))

	if err != nil {
		logger.Errorf("error marshalling response: %v", err)
		return app_errors.NewUnexpectedError(nil, nil)
	}

	_, err = w.Write(response)

	if err != nil {
		logger.Errorf("error writing response %v", err)
		return app_errors.NewUnexpectedError(nil, nil)
	}

	return nil
}
