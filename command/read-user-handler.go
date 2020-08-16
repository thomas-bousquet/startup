package command

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/thomas-bousquet/startup/api/adapter"
	. "github.com/thomas-bousquet/startup/repository"
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

func (h ReadUserCommand) Execute(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	user, err := h.userRepository.FindUser(id)

	if err != nil {
		return err
	}

	response, err := json.Marshal(adapter.NewUserAdapter(user))

	if err != nil {
		return err
	}
	_, err = w.Write(response)

	return err
}
