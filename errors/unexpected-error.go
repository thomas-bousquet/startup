package errors

import "net/http"

func NewUnexpectedError() CustomError {
	return  NewBaseError(http.StatusInternalServerError, "unexpected-error", http.StatusText(http.StatusInternalServerError), nil)
}
