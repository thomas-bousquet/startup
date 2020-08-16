package errors

import (
	"github.com/thomas-bousquet/startup/utils/validator"
)

type ValidationError struct {
	BaseError
	Errors []validator.Error `json:"errors"`
}

func NewValidationError(errors []validator.Error) ValidationError {
	return ValidationError{BaseError: BaseError{Message: "validation-error"}, Errors: errors}
}
