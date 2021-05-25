package error_handler

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/user-service/app_errors"
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

func appErrorToHttpResponse(error *app_errors.AppError) HttpError {
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

func (h ErrorHandler) WriteJSONErrorResponse(w http.ResponseWriter, e error, logger *logrus.Logger) {
	var appErrorType *app_errors.AppError
	if errors.As(e, &appErrorType) {
		httpError := appErrorToHttpResponse((e).(*app_errors.AppError))
		w.WriteHeader(httpError.Code)

		body, err := json.Marshal(e)

		if err != nil {
			logger.Errorf("e marshalling response: %v", err)
			http.Error(w, defaultErrorMessage, defaultErrorCode)
		}

		_, err = w.Write(body)

		if err != nil {
			logger.Errorf("e writing response: %v", err)
			http.Error(w, defaultErrorMessage, defaultErrorCode)
		}
	}
	// TODO: Handler other cases
}

func NewErrorHandler() ErrorHandler {
	return ErrorHandler{}
}
