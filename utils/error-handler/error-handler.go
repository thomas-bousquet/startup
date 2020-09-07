package error_handler

import (
	"encoding/json"
	"net/http"
	. "github.com/thomas-bousquet/startup/errors"
)

func WriteJSONErrorResponse(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case ValidationError:
		w.WriteHeader(e.HttpCode)
		doWriteError(w, err)
	case UnexpectedError:
		w.WriteHeader(e.HttpCode)
		doWriteError(w, err)
	case AuthorizationError:
		w.WriteHeader(e.HttpCode)
		doWriteError(w, err)
	default:
		doWriteError(w, NewUnexpectedError())
	}
}

func doWriteError(w http.ResponseWriter, error error) {
	defaultErrorMessage := http.StatusText(http.StatusInternalServerError)
	defaultHTTPCode := http.StatusInternalServerError

	body, err := json.Marshal(error)

	if err != nil {
		http.Error(w, defaultErrorMessage, defaultHTTPCode)
	}

	_, writeError := w.Write(body)

	if writeError != nil {
		http.Error(w, defaultErrorMessage, defaultHTTPCode)
	}
}
