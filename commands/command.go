package commands

import (
	"github.com/sirupsen/logrus"
	"net/http"
)

type Command interface {
	Execute(w http.ResponseWriter, r *http.Request, logger *logrus.Logger) error
}
