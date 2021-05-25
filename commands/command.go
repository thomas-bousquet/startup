package commands

import (
	"github.com/sirupsen/logrus"
	"github.com/thomas-bousquet/user-service/errors"
	"net/http"
)

type Command interface {
	Execute(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) *errors.AppError
}
