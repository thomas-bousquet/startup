package errors

import (
	"fmt"
)

type resource string

const (
	USER resource = "user"
)

func NewResourceNotFoundError(resource resource, id string) *AppError {
	message := fmt.Sprintf("%s with id %s was not found", resource, id)
	return NewError("not-found-error", &message, map[string]interface{}{"resource": resource})
}
