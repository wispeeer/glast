package logger

import (
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var timeFormatStr string = "15:04:05.000Z0700"

func init() {
	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: timeFormatStr,
		CallerPrettyfier: func(f *runtime.Frame) (function string, file string) {
			return "", fmt.Sprintf("%s:%d\t", path.Base(f.File), f.Line)
		},
	})
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
	log.SetLevel(logrus.DebugLevel)
}

func Task(task string) *logrus.Entry {
	return log.WithFields(logrus.Fields{
		"task": task,
	})
}
