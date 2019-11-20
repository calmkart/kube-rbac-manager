package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func InitLogger() {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		Log.Out = file
	} else {
		Log.Info("Failed to log to file")
	}
	Log.SetLevel(logrus.WarnLevel)
	Log.Formatter = &logrus.JSONFormatter{}
}
