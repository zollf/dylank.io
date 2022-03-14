package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var _log *logrus.Logger

func Log() *logrus.Logger {
	if _log == nil {
		_log = logrus.New()
		_log.SetFormatter(&logrus.TextFormatter{
			ForceColors:   true,
			FullTimestamp: true,
		})
		_log.SetOutput(os.Stdout)
	}
	return _log
}
