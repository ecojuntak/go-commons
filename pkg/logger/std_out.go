package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

func NewStdLogger(level logrus.Level, formatter logrus.Formatter) *logrus.Logger {
	return NewLogger(os.Stdout, level, formatter)
}

func JsonStdLogger(level logrus.Level) *logrus.Logger {
	return NewLogger(os.Stdout, level, &logrus.JSONFormatter{})
}

func TextStdLogger(level logrus.Level) *logrus.Logger {
	return NewLogger(os.Stdout, level, &logrus.TextFormatter{})
}
