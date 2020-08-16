package custom_error

import (
	"github.com/thomas-bousquet/startup/utils"
)

type ValidationError struct {
	Message string       `json:"message"`
	Errors []utils.Error `json:"errors"`
}

func (v ValidationError) Error() string {
	return v.Message
}

func NewValidationError(errors []utils.Error) ValidationError {
	return ValidationError{Errors: errors, Message: "validation-error"}
}


