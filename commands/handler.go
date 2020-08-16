package commands

import (
	"encoding/json"
	. "github.com/thomas-bousquet/startup/errors"
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
		switch e := err.(type) {
		case ValidationError:
			body, _ := json.Marshal(e)
			w.WriteHeader(http.StatusBadRequest)
			w.Write(body)
		case UnexpectedError:
			body, _ := json.Marshal(e)
			w.WriteHeader(http.StatusInternalServerError)
			w.Write(body)
		case AuthenticationError:
			body, _ := json.Marshal(e)
			w.WriteHeader(http.StatusUnauthorized)
			w.Write(body)
		default:
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}
}
