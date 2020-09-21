package logger

import (
	"os"
	"time"

	"github.com/sirupsen/logrus"
)

const PATH_LOG_FILE = "/tmp/"

type LoggerWrap struct {
	logrus.Logger
}

// GenerateNameLoc function for generate name and location
func GenerateNameLoc() string {
	dt := time.Now()
	return PATH_LOG_FILE + "icanvas-analytics-" + dt.Format("2006-01-02") + ".log"
}

// New init
func New() *LoggerWrap {
	baseLog := logrus.New()
	baseLog.SetReportCaller(true)

	baseLog.SetOutput(os.Stdout)

	logfile := GenerateNameLoc()

	file, err := os.OpenFile(logfile, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0755)
	if err != nil {
		baseLog.Fatal(err)
	}

	baseLog.SetOutput(file)

	return &LoggerWrap{
		*baseLog,
	}
}
