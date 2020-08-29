package errors

import "net/http"

type UnexpectedError struct {
	BaseError
}

func NewUnexpectedError() UnexpectedError {
	return UnexpectedError{BaseError: NewBaseError(http.StatusInternalServerError, "unexpected-error", "An unexpected error occurred")}
}
