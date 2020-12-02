package commands

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/startup/api/adapters"
	"github.com/thomas-bousquet/startup/errors"
	. "github.com/thomas-bousquet/startup/repositories"
	"net/http"
)

type ReadUserCommand struct {
	userRepository UserRepository
}

func NewReadUserCommand(userRepository UserRepository) ReadUserCommand {
	return ReadUserCommand{
		userRepository,
	}
}

func (c ReadUserCommand) Execute(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) *errors.Error {
	logger.Info("Reading user")

	vars := mux.Vars(r)
	id := vars["id"]

	user, err := c.userRepository.FindUser(id)

	if err != nil {
		logger.Errorf("error finding user by id: %v", err)
		return errors.NewUnexpectedError()
	}

	if user == nil {
		return errors.NewNotFoundError("user")
	}

	response, err := json.Marshal(adapters.NewUserAdapter(user))

	if err != nil {
		logger.Errorf("error marshalling response: %v", err)
		return errors.NewUnexpectedError()
	}

	_, err = w.Write(response)

	if err != nil {
		logger.Errorf("error writing response: %v", err)
		return errors.NewUnexpectedError()
	}

	return nil
}
