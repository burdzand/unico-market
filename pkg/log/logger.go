package log

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/sirupsen/logrus"
)

func init() {

	formatLogger, ok := os.LookupEnv("LOG_FORMAT")
	if !ok {
		formatLogger = "json"
	}
	formattimeLogger, ok := os.LookupEnv("LOG_TIMESTAMP_FORMAT")
	if !ok {
		formattimeLogger = "2006-01-02 15:04:05"
	}

	if formatLogger == "text" {
		textFormatter := logrus.TextFormatter{}
		textFormatter.FullTimestamp = true
		textFormatter.TimestampFormat = formattimeLogger
		logrus.SetFormatter(&textFormatter)
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)

	} else if formatLogger == "json" {
		jsonFormatter := logrus.JSONFormatter{}
		jsonFormatter.TimestampFormat = formattimeLogger
		logrus.SetFormatter(&jsonFormatter)
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)
	}

}

//Logger Struct
type Logger struct {
	*logrus.Entry
}

//NewLogger provider new instance of log
func NewLogger() *Logger {

	pc, pathfilename, _, ok := runtime.Caller(1)
	details := runtime.FuncForPC(pc)

	//var source string
	var filename string

	if ok && details != nil {
		//source = details.Name()
		//source = filepath.Base(source)
		_, filename = filepath.Split(pathfilename)
	} else {
		_, filename = filepath.Split(pathfilename)
	}

	_logger := logrus.WithFields(logrus.Fields{
		"filename": filename,
	})
	return &Logger{_logger}
}
