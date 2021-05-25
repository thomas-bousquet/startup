package errors

func NewUnexpectedError(message *string, metadata map[string]interface{}) *AppError {
	return NewError("unexpected-error", message, metadata)
}
