package errors

import "net/http"

func NewUnexpectedError() *Error {
	return NewError(http.StatusInternalServerError, "unexpected-error", http.StatusText(http.StatusInternalServerError), nil)
}
