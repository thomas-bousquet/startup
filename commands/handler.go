package commands

import (
	"github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/user-service/utils/error-handler"
	"net/http"
)

type Handler struct {
	command      Command
	logger       *logrus.Logger
	errorHandler ErrorHandler
}

func NewHandler(c Command, logger *logrus.Logger, errorHandler ErrorHandler) Handler {
	return Handler{command: c, logger: logger, errorHandler: errorHandler}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := h.command.Execute(w, r, h.logger)

	if err != nil {
		h.errorHandler.WriteJSONErrorResponse(w, err, h.logger)
	}
}
