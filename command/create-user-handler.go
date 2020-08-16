package command

import (
	"encoding/json"
	. "github.com/thomas-bousquet/startup/custom-error"
	. "github.com/thomas-bousquet/startup/model"
	. "github.com/thomas-bousquet/startup/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type CreateUserCommand struct {
	userRepository UserRepository
	validator      *validator.Validate
}

func NewCreateUserCommand(userRepository UserRepository, validator *validator.Validate) CreateUserCommand {
	return CreateUserCommand{
		userRepository,
		validator,
	}
}

func extractErrors(error error) []validator.FieldError {
	if error != nil {
		return error.(validator.ValidationErrors)
	} else {
		return nil
	}
}

func buildValidationErrors(errors []validator.FieldError, validationsErrors []ValidationErrorItem) []ValidationErrorItem {
	if len(errors) == 0 {
		return validationsErrors
	}
	nextError, remainingErrors := errors[0], errors[1:]

	validationError := ValidationErrorItem{
		Field:  nextError.Field(),
		Value:  nextError.Param(),
		Reason: nextError.Tag(),
	}
	return buildValidationErrors(remainingErrors, append(validationsErrors, validationError))
}

func getValidationErrors(error error) []ValidationErrorItem {
	validationErrors := extractErrors(error)
	return buildValidationErrors(validationErrors, nil)
}

func (h CreateUserCommand) Execute(w http.ResponseWriter, r *http.Request) error {
	var user = User{}
	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		return err
	}

	err = h.validator.Struct(user)
	errors := getValidationErrors(err)

	if len(errors) > 0 {
		return NewValidationError(errors)
	}

	userId, err := h.userRepository.CreateUser(user)

	if err != nil {
		return err
	}

	response, err := json.Marshal(map[string]primitive.ObjectID{"id": userId})

	if err != nil {
		return err
	}

	w.WriteHeader(http.StatusCreated)
	 _, err = w.Write(response)

	return err
}
