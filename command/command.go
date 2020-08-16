package command

import (
	"net/http"
)

type Command interface {
	Execute(w http.ResponseWriter, r *http.Request) error
}
