package errors

import "net/http"

func NewAuthorizationError(message string) *Error {
	return NewError(http.StatusUnauthorized, "authorization-error", message, nil)
}
