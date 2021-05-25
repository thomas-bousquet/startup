package commands

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/user-service/api/adapters"
	"github.com/thomas-bousquet/user-service/errors"
	. "github.com/thomas-bousquet/user-service/repositories"
	"net/http"
)

type ReadUsersCommand struct {
	userRepository UserRepository
}

func NewReadUsersCommand(userRepository UserRepository) ReadUsersCommand {
	return ReadUsersCommand{
		userRepository,
	}
}

func (c ReadUsersCommand) Execute(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) *errors.AppError {
	users, err := c.userRepository.FindUsers()

	if err != nil {
		logger.Errorf("error finding users: %v", err)
		return errors.NewUnexpectedError(nil, nil)
	}

	var usersAdapter []adapters.UserAdapter

	for _, user := range users {
		userAdapter := adapters.NewUserAdapter(&user)
		usersAdapter = append(usersAdapter, userAdapter)
	}

	response, err := json.Marshal(usersAdapter)

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
