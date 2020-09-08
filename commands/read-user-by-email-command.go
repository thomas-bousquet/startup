package commands

import (
	"encoding/json"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/startup/api/adapters"
	. "github.com/thomas-bousquet/startup/repositories"
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

func (c ReadUserByEmailCommand) Execute(w http.ResponseWriter, r *http.Request) error {
	log.Info("Read user by email")
	vars := mux.Vars(r)
	email := vars["email"]

	user, err := c.userRepository.FindUserByEmail(email)

	if err != nil {
		return err
	}

	response, err := json.Marshal(adapters.NewUserAdapter(user))

	if err != nil {
		return err
	}

	_, err = w.Write(response)

	if err != nil {
		return err
	}

	return nil
}