package error_handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/startup/errors"
	"net/http"
)

type ErrorHandler struct{}

func (h ErrorHandler) WriteJSONErrorResponse(w http.ResponseWriter, error *errors.Error, logger *logrus.Logger) {
	w.WriteHeader(error.HttpCode)

	defaultErrorMessage := http.StatusText(http.StatusInternalServerError)
	defaultHTTPCode := http.StatusInternalServerError

	body, err := json.Marshal(error)

	if err != nil {
		logger.Errorf("error marshalling response: %v", err)
		http.Error(w, defaultErrorMessage, defaultHTTPCode)
	}

	_, err = w.Write(body)

	if err != nil {
		logger.Errorf("error writing response: %v", err)
		http.Error(w, defaultErrorMessage, defaultHTTPCode)
	}
}

func NewErrorHandler() ErrorHandler {
	return ErrorHandler{}
}
