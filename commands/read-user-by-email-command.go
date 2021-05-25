package commands

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/user-service/api/adapters"
	"github.com/thomas-bousquet/user-service/errors"
	. "github.com/thomas-bousquet/user-service/repositories"
	"net/http"
)

type ReadUserByEmailCommand struct {
	userRepository UserRepository
}

func NewReadUserByEmailCommand(userRepository UserRepository) ReadUserByEmailCommand {
	return ReadUserByEmailCommand{
		userRepository,
	}
}

func (c ReadUserByEmailCommand) Execute(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) *errors.AppError {
	vars := mux.Vars(r)
	email := vars["email"]

	user, err := c.userRepository.FindUserByEmail(email)

	if err != nil {
		logger.Errorf("%v", err)
		return errors.NewUnexpectedError(nil, nil)
	}

	if user == nil {
		return errors.NewResourceNotFoundError(errors.USER, email)
	}

	response, err := json.Marshal(adapters.NewUserAdapter(user))

	if err != nil {
		logger.Errorf("error marshalling response: %v", err)
		return errors.NewUnexpectedError(nil, nil)
	}

	_, err = w.Write(response)

	if err != nil {
		logger.Errorf("error writing response: %v", err)
		return errors.NewUnexpectedError(nil, nil)
	}

	return nil
}
