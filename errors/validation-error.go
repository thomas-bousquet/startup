package errors

import (
	"startup/utils/validator"
	"net/http"
)

type ValidationError struct {
	BaseError
	Errors []validator.Error `json:"errors"`
}

func NewValidationError(message string, errors []validator.Error) ValidationError {
	return ValidationError{BaseError: NewBaseError(http.StatusBadRequest, "validation-error", message), Errors: errors}
}
