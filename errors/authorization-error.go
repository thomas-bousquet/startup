package errors

import "net/http"

type AuthorizationError struct {
	BaseError
}

func NewAuthorizationError(message string) AuthorizationError {
	return AuthorizationError{BaseError: NewBaseError(http.StatusUnauthorized, "authorization-error", message)}
}
