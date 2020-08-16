package validator

import (
	"gopkg.in/go-playground/validator.v9"
	"strings"
)

type Validator struct {
	validator *validator.Validate
}

func (v Validator) ValidateStruct(data interface{}) []Error {
	err := v.validator.Struct(data)
	validationErrors := extractErrors(err)
	return buildValidationErrors(validationErrors, nil)
}

func (v Validator) ValidateStructExcept(data interface{}, fields ...string) []Error {
	err := v.validator.StructExcept(data, fields...)
	validationErrors := extractErrors(err)
	return buildValidationErrors(validationErrors, nil)
}

func NewValidator(v *validator.Validate) Validator {
	return Validator{validator: v}
}

func buildValidationErrors(errs []validator.FieldError, validationsErrors []Error) []Error {
	if len(errs) == 0 {
		return validationsErrors
	}
	nextError, remainingErrors := errs[0], errs[1:]

	validationError := Error{
		Field:  strings.ToLower(nextError.Field()),
		Value:  strings.ToLower(nextError.Param()),
		Reason: nextError.Tag(),
	}
	return buildValidationErrors(remainingErrors, append(validationsErrors, validationError))
}

func extractErrors(error error) []validator.FieldError {
	if error != nil {
		return error.(validator.ValidationErrors)
	} else {
		return nil
	}
}

type Error struct {
	Field  string `json:"field"`
	Value  string `json:"value"`
	Reason string `json:"reason"`
}
