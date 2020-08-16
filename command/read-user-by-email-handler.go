package command

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/thomas-bousquet/startup/api/adapter"
	. "github.com/thomas-bousquet/startup/repository"
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

func (h ReadUserByEmailCommand) Execute(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	email := vars["email"]

	user, err := h.userRepository.FindUserByEmail(email)

	if err != nil {
		return err
	}

	response, err := json.Marshal(adapter.NewUserAdapter(user))

	if err != nil {
		return err
	}

	_, err = w.Write(response)

	if err != nil {
		return err
	}

	return nil
}
