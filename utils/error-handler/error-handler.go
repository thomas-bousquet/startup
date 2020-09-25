package error_handler

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"net/http"
	. "startup/errors"
)

func WriteJSONErrorResponse(w http.ResponseWriter, err error) {
	switch e := err.(type) {
	case ValidationError:
		log.Error(e)
		w.WriteHeader(e.HttpCode)
		doWriteError(w, err)
	case UnexpectedError:
		log.Error(e)
		w.WriteHeader(e.HttpCode)
		doWriteError(w, err)
	case AuthorizationError:
		log.Error(e)
		w.WriteHeader(e.HttpCode)
		doWriteError(w, err)
	default:
		unexpectedError :=  NewUnexpectedError()
		log.Error(unexpectedError)
		doWriteError(w, unexpectedError)
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
