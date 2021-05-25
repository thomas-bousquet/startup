package app_errors

import (
	"github.com/thomas-bousquet/user-service/utils/validator"
)

func NewValidationError(message string, errors []validator.Error) error {
	return NewAppError("validation-error", &message, map[string]interface{}{"validation-errors": errors})
}
