package commands

import (
	errorHandler "startup/utils/error-handler"
	"net/http"
)

type Handler struct {
	Command Command
}

func NewHandler(c Command) Handler {
	return Handler{Command: c}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	err = h.Command.Execute(w, r)

	if err != nil {
		errorHandler.WriteJSONErrorResponse(w, err)
	}
}
