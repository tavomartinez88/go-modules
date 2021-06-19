package interfaces

import (
	"github.com/sirupsen/logrus"
	"time"
)

type ILogger interface {
	GetLogger(time time.Time) *logrus.Logger
}