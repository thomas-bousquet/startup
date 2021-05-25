package app_errors

func NewUnexpectedError(message *string, metadata map[string]interface{}) error {
	return NewAppError("unexpected-error", message, metadata)
}
