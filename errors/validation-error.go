package errors

import (
	"github.com/thomas-bousquet/startup/utils/validator"
	"net/http"
)

func NewValidationError(message string, errors []validator.Error) CustomError {
 return NewBaseError(http.StatusBadRequest, "validation-error", message, map[string][]validator.Error{"validation": errors})
}