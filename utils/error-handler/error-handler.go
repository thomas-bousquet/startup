package error_handler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/user-service/errors"
	"net/http"
)

type ErrorHandler struct{}

type HttpResponse struct {
	HttpCode int
	Key      string
	Message  string
	Metadata map[string]interface{}
}

var errorHttpCodes = map[string]int{
	"authorization-error": http.StatusUnauthorized,
	"validation-error":    http.StatusBadRequest,
	"not-found-error":     http.StatusNotFound,
	"unexpected-error":    http.StatusInternalServerError,
}

type HttpError struct {
	Code    int    `json:"-"`
	Key     string `json:"key"`
	Message string `json:"message,omitempty"`
}

var defaultErrorKey = "unexpected-error"
var defaultErrorMessage = ""
var defaultErrorCode = http.StatusInternalServerError

func appErrorToHttpResponse(error *errors.AppError) HttpError {
	httpKey := defaultErrorKey
	httpMessage := defaultErrorMessage
	httpCode := defaultErrorCode

	_, exists := errorHttpCodes[error.Key]

	if exists {
		httpKey = error.Key
		httpCode = errorHttpCodes[error.Key]
		if error.Message != nil {
			httpMessage = *(error.Message)
		} else {
			httpMessage = ""
		}
	}

	return HttpError{
		Code:    httpCode,
		Key:     httpKey,
		Message: httpMessage,
	}
}

func (h ErrorHandler) WriteJSONErrorResponse(w http.ResponseWriter, error *errors.AppError, logger *logrus.Logger) {

	httpError := appErrorToHttpResponse(error)
	w.WriteHeader(httpError.Code)

	body, err := json.Marshal(error)

	if err != nil {
		logger.Errorf("error marshalling response: %v", err)
		http.Error(w, defaultErrorMessage, defaultErrorCode)
	}

	_, err = w.Write(body)

	if err != nil {
		logger.Errorf("error writing response: %v", err)
		http.Error(w, defaultErrorMessage, defaultErrorCode)
	}
}

func NewErrorHandler() ErrorHandler {
	return ErrorHandler{}
}
