package errors

import (
	"fmt"
	"net/http"
)

func NewNotFoundError(entity string) *Error {
	return NewError(http.StatusUnauthorized, "not-found", fmt.Sprintf("%s was not found", entity), nil)
}
