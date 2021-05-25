package errors

func NewAuthorizationError(metadata map[string]interface{}) *AppError {
	return NewError("authorization-error", nil, metadata)
}
