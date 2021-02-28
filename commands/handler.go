package commands

import (
	"github.com/sirupsen/logrus"
	. "github.com/thomas-bousquet/user-service/utils/error-handler"
	"net/http"
	"reflect"
)

type Handler struct {
	command      Command
	logger       *logrus.Logger
	errorHandler ErrorHandler
	commandName  string
}

func NewHandler(c Command, logger *logrus.Logger, errorHandler ErrorHandler) Handler {
	return Handler{command: c, logger: logger, errorHandler: errorHandler, commandName: getCommandName(c)}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Start", h.commandName)
	err := h.command.Execute(w, r, h.logger)
	h.logger.Info("Done", h.commandName)

	if err != nil {
		h.errorHandler.WriteJSONErrorResponse(w, err, h.logger)
	}
}

func getCommandName(i interface{}) string {
	var name string

	if t := reflect.TypeOf(i); t.Kind() == reflect.Ptr {
		name = t.Elem().Name()
	} else {
		name = t.Name()
	}

	return name
}
