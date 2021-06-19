package loggerImpl

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

type loggerLogrus struct{
	logger *logrus.Logger
}

func (o loggerLogrus) GetLogger(time time.Time) *logrus.Logger {
	if o.logger == nil {
		o.logger = new(logrus.Logger)
		o.logger.SetFormatter(&logrus.JSONFormatter{})
		o.logger.SetOutput(os.Stdout)
		o.logger.SetLevel(logrus.DebugLevel)
	}

	o.logger.WithTime(time)

	return o.logger
}

