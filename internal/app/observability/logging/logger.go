package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

func NewLogger(level string) *logrus.Logger {
	logger := logrus.New()
	logger.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02T15:04:05Z07:00",
	})
	logger.SetOutput(os.Stdout)

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.InfoLevel
	}
	logger.SetLevel(lvl)
	return logger
}
