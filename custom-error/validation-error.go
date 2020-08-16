package custom_error

type ValidationError struct {
	Message string `json:"message"`
	Errors []ValidationErrorItem `json:"errors"`
}

func (v ValidationError) Error() string {
	return v.Message
}

func NewValidationError(errors []ValidationErrorItem) ValidationError {
	return ValidationError{Errors: errors, Message: "validation-error"}
}

type ValidationErrorItem struct {
	Field string `json:"field"`
	Value string `json:"value"`
	Reason string `json:"reason"`
}
