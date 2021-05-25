package errors

import (
	"github.com/thomas-bousquet/user-service/utils/validator"
)

func NewValidationError(message string, errors []validator.Error) *AppError {
	return NewError("validation-error", &message, map[string]interface{}{"validation-errors": errors})
}
