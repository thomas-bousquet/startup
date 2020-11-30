package commands

import (
	"github.com/sirupsen/logrus"
	errorHandler "github.com/thomas-bousquet/startup/utils/error-handler"
	"net/http"
)

type Handler struct {
	Command Command
	Logger *logrus.Logger
}

func NewHandler(c Command, logger *logrus.Logger) Handler {
	return Handler{Command: c, Logger: logger}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var err error
	err = h.Command.Execute(w, r, h.Logger)

	if err != nil {
		errorHandler.WriteJSONErrorResponse(w, err, h.Logger)
	}
}
