package log

import (
	"os"
	"github.com/Prasanna-18/temp/api/configu"
	"github.com/sirupsen/logrus"
)

var (
	Log *logrus.Logger
)


func init() {
	level,err:= logrus.ParseLevel(configu.Loglevel)
	if err!= nil {
		level = logrus.DebugLevel
	}
	Log = &logrus.Logger{
		Level: level,
		Out: os.Stdout,
	}
if configu.IsProductionn() {
Log.Formatter =&logrus.JSONFormatter{}
}else{
Log.Formatter= &logrus.TextFormatter{}
}

}
