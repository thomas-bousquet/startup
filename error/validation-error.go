package error

import (
	"errors"
	"net/http"
)

type ValidationError struct {
	RequestError `json:"-"`
	Errors []ValidationErrorItem `json:"errors"`
}

func (v ValidationError) Error() error {
	return errors.New("invalid-request")
}

func NewValidationError(errors []ValidationErrorItem) ValidationError {
	return ValidationError{RequestError: RequestError{StatusCode: http.StatusBadRequest}, Errors: errors}
}
