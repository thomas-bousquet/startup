package commands

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/user-service/api/adapters"
	"github.com/thomas-bousquet/user-service/app_errors"
	. "github.com/thomas-bousquet/user-service/repositories"
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

func (c ReadUserCommand) Execute(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) error {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := c.userRepository.FindUser(id)

	if err != nil {
		logger.Errorf("error finding user by id %q: %v", id, err)
		return app_errors.NewUnexpectedError(nil, nil)
	}

	if user == nil {
		return app_errors.NewResourceNotFoundError(app_errors.USER, id)
	}

	response, err := json.Marshal(adapters.NewUserAdapter(user))

	if err != nil {
		logger.Errorf("error marshalling response: %v", err)
		return app_errors.NewUnexpectedError(nil, nil)
	}

	_, err = w.Write(response)

	if err != nil {
		logger.Errorf("error writing response: %v", err)
		return app_errors.NewUnexpectedError(nil, nil)
	}

	return nil
}
