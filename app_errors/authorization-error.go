package app_errors

func NewAuthorizationError(metadata map[string]interface{}) error {
	return NewAppError("authorization-error", nil, metadata)
}
