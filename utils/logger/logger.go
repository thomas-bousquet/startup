package logger

import (
	"github.com/sirupsen/logrus"
)

func NewLogger() *logrus.Logger {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	logger.SetReportCaller(true)
	logger.SetFormatter(&logrus.JSONFormatter{})
	return logger
}
