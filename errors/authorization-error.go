package errors

import "net/http"

func NewAuthorizationError(message string) CustomError {
	return NewBaseError(http.StatusUnauthorized, "authorization-error", message, nil)
}
