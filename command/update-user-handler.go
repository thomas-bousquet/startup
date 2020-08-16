package command

import (
	"encoding/json"
	"github.com/gorilla/mux"
	. "github.com/thomas-bousquet/startup/model"
	. "github.com/thomas-bousquet/startup/repository"
	"github.com/thomas-bousquet/startup/utils"
	"net/http"
)

type UpdateUserCommand struct {
	userRepository UserRepository
	validator      utils.Validator
}

func NewUpdateUserCommand(userRepository UserRepository, validator utils.Validator) UpdateUserCommand {
	return UpdateUserCommand{
		userRepository,
		validator,
	}
}

func (h UpdateUserCommand) Execute(w http.ResponseWriter, r *http.Request) error {
	vars := mux.Vars(r)
	id := vars["id"]

	user := User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		return err
	}

	user.Id = id

	// TODO: Implement ValidateStructExcept in utils/Validator then use it right here
	//err = h.validator.StructExcept(user, "password")
	//for _, e := range err.(validator.ValidationErrors) {
	//	fmt.Println(e)
	//}

	err = h.userRepository.UpdateUser(id, user)

	if err != nil {
		return err
	}

	return nil
}