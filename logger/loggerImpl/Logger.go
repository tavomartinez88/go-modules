package loggerImpl

import (
	"github.com/sirupsen/logrus"
	"time"
)

//GetLogger return a pointer to *logrus.Logger, ready to log events and other tipics
func GetLogger() *logrus.Logger {
	return loggerLogrus{}.GetLogger(time.Now())
}
