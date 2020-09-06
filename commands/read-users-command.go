package commands

import (
	"encoding/json"
	"github.com/thomas-bousquet/startup/api/adapters"
	. "github.com/thomas-bousquet/startup/repositories"
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

func (c ReadUsersCommand) Execute(w http.ResponseWriter, r *http.Request) error {
	users, err := c.userRepository.FindUsers()

	if err != nil {
		return err
	}

	var usersAdapter []adapters.UserAdapter

	for _, user := range users {
		userAdapter := adapters.NewUserAdapter(&user)
		usersAdapter = append(usersAdapter, userAdapter)
	}

	response, err := json.Marshal(usersAdapter)
	_, err = w.Write(response)

	return err
}
