package handler

import (
	"encoding/json"
	. "github.com/thomas-bousquet/startup/error"
	. "github.com/thomas-bousquet/startup/model"
	. "github.com/thomas-bousquet/startup/repository"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gopkg.in/go-playground/validator.v9"
	"net/http"
)

type CreateUserHandler struct {
	userRepository UserRepository
	validator      *validator.Validate
}

func NewCreateUserHandler(userRepository UserRepository, validator *validator.Validate) CreateUserHandler {
	return CreateUserHandler{
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

func (h CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var user = User{}
	json.NewDecoder(r.Body).Decode(&user)

	err := h.validator.Struct(user)
	errors := getValidationErrors(err)

	if len(errors) > 0 {
		validationError := NewValidationError(errors)
		body, _ := json.Marshal(validationError)
		w.WriteHeader(validationError.StatusCode)
		w.Write(body)
		return
	}

	userId := h.userRepository.CreateUser(user)

	response, _ := json.Marshal(map[string]primitive.ObjectID{"id": userId})
	w.WriteHeader(http.StatusCreated)
	w.Write(response)
}
