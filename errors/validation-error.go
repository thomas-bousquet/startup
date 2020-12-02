package errors

import (
	"github.com/thomas-bousquet/startup/utils/validator"
	"net/http"
)

func NewValidationError(message string, errors []validator.Error) *Error {
 return NewError(http.StatusBadRequest, "validation-error", message, map[string][]validator.Error{"validation": errors})
}