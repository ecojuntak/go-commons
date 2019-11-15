package logger

import (
	"github.com/sirupsen/logrus"
	"io"
)

func NewLogger(writer io.Writer, level logrus.Level, formatter logrus.Formatter) *logrus.Logger {
	logger := logrus.New()
	logger.SetOutput(writer)
	logger.SetLevel(level)
	if formatter != nil {
		logger.SetFormatter(formatter)
	}

	return logger
}
