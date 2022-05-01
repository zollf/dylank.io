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
			DisableColors: true,
			ForceColors:   false,
			FullTimestamp: true,
		})
		_log.SetOutput(os.Stdout)
	}
	return _log
}
